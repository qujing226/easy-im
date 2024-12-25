import Vue from 'vue'
import App from './App'
import ElementUI from 'element-ui'
import LemonIMUI from 'lemon-imui';
import 'lemon-imui/dist/index.css';
import '@/permission'
import router from './router'
import store from './store'
import Api from "@/api/index.js"
import Clipboard from 'v-clipboard'
Vue.use(ElementUI, { size: 'small' })
Vue.use(LemonIMUI);
Vue.use(store);
Vue.use(Clipboard)
Vue.prototype.$api = Api
new Vue({
    el: '#app',
    router,
    store,
    components: { App },
    template: '<App/>'
})