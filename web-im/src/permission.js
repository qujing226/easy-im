import router from './router'
import {
    getAuth
} from '@/utils/auth' 
import lockr from 'lockr'
const whiteList = ['/login','/register'] 
const closeList=['/','/chat'] 
router.beforeEach((to, from, next) => {
    if (getAuth()) {
		console.log("true");
        let config=lockr.get('globalConfig');
        let demon=config.demon_mode;
        let toPath='';
        let userInfo=lockr.get('UserInfo');
        if (whiteList.includes(to.path) || (to.path=='/' && toPath)) {
            next({
                path: toPath
            })
		}else {
            next()
        }
    } else {
		console.log("false");
        if (whiteList.indexOf(to.path) !== -1) {
            next()
        } else {
            next(`/login`)
        }
    }
})