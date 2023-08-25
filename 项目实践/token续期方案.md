## token续期方案
- 刷新Token的过期时间：在Token的过期时间到期前，可以通过向认证服务器发起请求，获取新的Token并刷新原Token的过期时间。这种方案需要确保Token的安全性，避免Token被恶意盗用。

- 使用Refresh Token：Refresh Token是一种特殊的Token，用于获取新的Access Token。当Access Token过期时，客户端可以使用Refresh Token向认证服务器请求新的Access Token。Refresh Token通常具有更长的有效期，但也需要保证其安全性。

- 自动续期：可以在客户端中设置一个定时器，在Token即将过期时自动向认证服务器请求新的Token。这种方案需要考虑网络延迟等因素，确保新Token在旧Token过期前就已经获取到。

这里使用的是 __刷新token的过期时间__ 这种方案

### 后端实现
### JWT中间件
实现方法：通过获取老的token ，再计算token剩余时间剩余时间，如果达到符合的时间将重新颁发新的token。
```go
const Authorization = "Authorization"

func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get(Authorization)
		if authHeader == "" {
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}
		// 按空格分隔
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		// parts[1]是获取到tokenString, 我们使用之前定义好的解析JWT的函数来解析它
		mc, err := jwt.ParseToken(parts[1])
		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

    // 重新颁发 token
		newToken, err, ok := jwt.RefreshToken(parts[1], mc.Username, mc.UserID )
		// 返回请求头 token
		if ok {
			c.Header(Authorization, newToken)
		}
		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		// 将当前请求的userID信息保存到请求的上下文上
		c.Set(controller.CtxUserIDKey, mc.UserID)
		c.Next() // 后续的处理请求的函数中，可以用c.Get(CtxUserIDKey)来获取当前请求的用户信息
	}
}
```
使用 __jwt.RefreshToken(parts[1], mc.Username, mc.UserID )__ 校验和生成新的token，再把新的token写入到响应头中，前端从响应头中获取新的token。
这里使用的是自定义状态来返回错误的信息，也可以使用 __c.AbortWithStatus(http.StatusInternalServerError)__ 来返回前端一个http的401状态码。

### token续期函数
```go
// RefreshToken token 续期
func RefreshToken(tokenString, username string, userID int64) (string, error, bool) {
	// 解析原始的 Token
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySecret, nil
	})

	if err != nil {
		return "", err, false
	}

	// 检查原始 Token 是否有效
	if !token.Valid {
		return "", errors.New("invalid token"), false
	}

	// 获取原始 Token 中的 claims 信息
	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", errors.New("failed to parse claims"), false
	}

	// 如果Token还有30分钟过期，则生成新的Token
	timeRemaining := time.Until(time.Unix(claims.ExpiresAt, 0))
	if timeRemaining <= 30 * time.Minute {
		newToken, err := GenToken(userID, username)
		if err != nil {
			return "", err, false
		}
		return newToken, nil, true
	}

	return "", nil, false
}
```
这里接收的参数可以根据自己的需要来进行修改，我这里需要一个老的token、用户名和用户的id。如果token剩余时间小于30分钟将生成一个新的token，并抛出一个错误，在jwt中间件中捕获到这个错误，返回一个信号将新token设置到响应头中。

## 前端实现
### 请求处理
在响应中获取到返回的data.code，如果等于设置的错误信息状态码即可进行处理。因为token错误发送的请求都会返回错误，所以避免用户token还没过期因为切换token而无法显示内容，这里实现请求重发并且重发次数在一定数量内，以免无限重发请求。
```typescript
if (data.code === ResultEnum.OVERDUE) {
	tryHideFullScreenLoading();
	store.dispatch(setToken(headers.authorization));
	if (headers.authorization) {
		config.headers!.Authorization = "";
		const retryCount = axiosCanceler.reset(config);
		if (retryCount >= 5) {
			axiosCanceler.removePending(config);
			return Promise.reject(data);
		}
		return this.service(config);
	}
	message.error(data.msg);
	window.location.href = "/admin";
	axiosCanceler.removePending(config);
	return Promise.reject(data);
}
```
完整代码
```typescript
class RequestHttp {
	service: AxiosInstance;
	public constructor(config: AxiosRequestConfig) {
		// 实例化axios
		this.service = axios.create(config);

		/**
		 * @description 请求拦截器
		 * 客户端发送请求 -> [请求拦截器] -> 服务器
		 * token校验(JWT) : 接受服务器返回的token,存储到redux/本地储存当中
		 */
		this.service.interceptors.request.use(
			(config: AxiosRequestConfig) => {
				NProgress.start();
				// * 将当前请求添加到 pending 中
				axiosCanceler.addPending(config);
				// * 如果当前请求不需要显示 loading,在api服务中通过指定的第三个参数: { headers: { noLoading: true } }来控制不显示loading，参见loginApi
				config.headers!.noLoading || showFullScreenLoading();
				const token: string = store.getState().global.token;
				return { ...config, headers: { ...config.headers, Authorization: `Bearer ${token}` } };
			},
			(error: AxiosError) => {
				return Promise.reject(error);
			}
		);

		/**
		 * @description 响应拦截器
		 *  服务器换返回信息 -> [拦截统一处理] -> 客户端JS获取到信息
		 */
		this.service.interceptors.response.use(
			(response: AxiosResponse) => {
				const { data, config, headers } = response;
				NProgress.done();
				if (data.code === ResultEnum.OVERDUE) {
					tryHideFullScreenLoading();
					store.dispatch(setToken(headers.authorization));
					if (headers.authorization) {
						config.headers!.Authorization = "";
						const retryCount = axiosCanceler.reset(config);
						if (retryCount >= 5) {
							axiosCanceler.removePending(config);
							return Promise.reject(data);
						}
						return this.service(config);
					}
					message.error(data.msg);
					window.location.href = "/admin";
					axiosCanceler.removePending(config);
					return Promise.reject(data);
				}
				// * 在请求结束后，移除本次请求(关闭loading)
				axiosCanceler.removePending(config);
				tryHideFullScreenLoading();
				// * token 续期失效
				// * 全局错误信息拦截（防止下载文件得时候返回数据流，没有code，直接报错）
				if (data.code && data.code !== ResultEnum.SUCCESS) {
					message.error(data.msg);
					return Promise.reject(data);
				}
				// * 成功请求（在页面上除非特殊情况，否则不用处理失败逻辑）
				return data;
			},
			async (error: AxiosError) => {
				const { response } = error;
				NProgress.done();
				tryHideFullScreenLoading();
				if (error.message === "canceled") return;
				// 请求超时单独判断，请求超时没有 response
				if (error.message.indexOf("timeout") !== -1) message.error("请求超时，请稍后再试");
				// 根据响应的错误状态码，做不同的处理
				if (response) checkStatus(response.status);
				// 服务器结果都没有返回(可能服务器错误可能客户端断网) 断网处理:可以跳转到断网页面
				if (!window.navigator.onLine) window.location.hash = "/500";
				return Promise.reject(error);
			}
		);
	}

	// * 常用请求方法封装
	get<T>(url: string, params?: object, _object = {}): Promise<ResultData<T>> {
		return this.service.get(url, { params, ..._object });
	}
	post<T>(url: string, params?: object, _object = {}): Promise<ResultData<T>> {
		return this.service.post(url, params, _object);
	}
	put<T>(url: string, params?: object, _object = {}): Promise<ResultData<T>> {
		return this.service.put(url, params, _object);
	}
	delete<T>(url: string, params?: any, _object = {}): Promise<ResultData<T>> {
		return this.service.delete(url, { params, ..._object });
	}
}
```
### 请求阻止
```typescript
import axios, { AxiosRequestConfig, Canceler } from "axios";
import { isFunction } from "@/utils/is/index";
import qs from "qs";

// * 声明一个 Map 用于存储每个请求的标识 和 取消函数
let pendingMap = new Map<string, Canceler>();
let retryMap = new Map<string, number>();

// * 序列化参数
export const getPendingUrl = (config: AxiosRequestConfig) =>
	[config.method, config.url, qs.stringify(config.data), qs.stringify(config.params)].join("&");

export class AxiosCanceler {
	/**
	 * @description: 添加请求
	 * @param {Object} config
	 */
	addPending(config: AxiosRequestConfig) {
		// * 在请求开始前，对之前的请求做检查取消操作
		this.removePending(config);
		const url = getPendingUrl(config);
		config.cancelToken =
			config.cancelToken ||
			new axios.CancelToken(cancel => {
				if (!pendingMap.has(url) || !retryMap.has(url)) {
					// 如果 pending 中不存在当前请求，则添加进去
					pendingMap.set(url, cancel);
				}
			});
	}

	/**
	 * @description: 移除请求
	 * @param {Object} config
	 */
	removePending(config: AxiosRequestConfig) {
		const url = getPendingUrl(config);

		if (pendingMap.has(url)) {
			// 如果在 pending 中存在当前请求标识，需要取消当前请求，并且移除
			const cancel = pendingMap.get(url);
			cancel && cancel();
			pendingMap.delete(url);
			retryMap.delete(url);
		}
	}

	/**
	 * @description: 清空所有pending
	 */
	removeAllPending() {
		pendingMap.forEach(cancel => {
			cancel && isFunction(cancel) && cancel();
		});
		pendingMap.clear();
		retryMap.clear();
	}

	/**
	 * @description: 重置
	 */
	reset(config: AxiosRequestConfig): number {
		const url = getPendingUrl(config);
		const retryCount = retryMap.get(url) || 0;
		retryMap.set(url, retryCount + 1);
		pendingMap = new Map<string, Canceler>();
		return retryCount;
	}
}
```