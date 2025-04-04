import axios from 'axios'

window.BASE_URL = "/api";
console.log(window.BASE_URL);
const service = axios.create({
    baseURL: window.BASE_URL, 
    timeout: 60000 
})

service.baseURL=window.BASE_URL;
// response 拦截器
service.interceptors.response.use(
    response => {
		return response.data
    },
)
 
export default service