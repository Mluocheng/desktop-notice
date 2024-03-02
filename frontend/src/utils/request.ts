import axios, { AxiosRequestConfig } from 'axios';

export function request(
	url: any,
	options?: string
): Promise<{ headers: { [s: string]: string }; data: any; config: any; status: number; statusText: string }> {
	const requestId = String(Math.random()).replace(/\./g, '').padStart(6, '0').substring(0, 6);
	console.log('[' + requestId + '] request ' + url + ' ' + JSON.stringify(options));

	// 先给属性名增加双引号，然后将值的单引号转换为双引号
	const correctedJsonString = options?.replace(/(\w+):/g, '"$1":').replace(/'/g, '"'); // 将单引号转换为双引号
	let _options: AxiosRequestConfig = {};
	// 现在尝试转换为JSON对象
	try {
	_options = correctedJsonString ? JSON.parse(correctedJsonString) : '';
	console.log(_options);
	} catch (e) {
	console.error("Parsing error:", e);
	}

	// 请求配置
	const axiosConfig = {
		method: 'get', // 可以是 'get', 'post', 'put', 'delete' 等
		url: url,
		..._options,
		headers: {
			'Content-Type': 'application/json',
			..._options?.headers
		},
		timeout: 1000 * 60 * 10
	};

	const responseRes = {
		headers: {},
		data: null,
		config: {},
		status: 0,
		statusText: ''
	};
	// 发送请求
	return axios(axiosConfig)
		.then((response: any /* { headers, data, config, status, statusText } */) => {
			console.log('[' + requestId + '] response:', JSON.stringify(response.data)); // 输出响应数据
			responseRes.headers = response.headers;
			responseRes.data = response.data;
			// responseRes.config = config;
			responseRes.status = response.status;
			responseRes.statusText = response.statusText;
			// console.log("响应", responseRes);
			return responseRes;
		})
		.catch((error: { message: any; response: { status: any; statusText: any; headers: any; }; }) => {
			return Promise.reject({
				message: error.message,
				response: {
					status: error.response?.status || 0,
					statusText: error.response?.statusText,
					// headers: JSON.parse(JSON.stringify(error.response?.headers))
				}
			});
		});
}
