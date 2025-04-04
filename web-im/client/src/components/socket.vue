<template>
    <div>
        <audio id="chatAudio">
        <source src="../assets/voice/notify.ogg" type="audio/ogg">
        <source src="../assets/voice/notify.mp3" type="audio/mpeg">
        <source src="../assets/voice/notify.wav" type="audio/wav">
        </audio>
    </div>
</template>

<script>
    import Vue from 'vue'
    import Lockr from 'lockr'

    export default {
        name: "socket",
        data() {
            return {
                websocket: null,
            }
        },
        methods: {
            getWsUrl(){
                return "ws://"+process.env.VUE_APP_BASE_WS+"/ws";
            },
            initWebSocket() { //初始化weosocket
                //ws地址
                const WS_URI = this.getWsUrl();
				let token =Lockr.get('authToken')
				this.websocket = new WebSocket(WS_URI,[token]);
				
				// 打开websocket
				this.websocket.onopen = event=> {
					
					console.log("成功连接");
					// this.websocketSend("连接");
				}
				
				//监听消息
				this.websocket.onmessage =e=>  {
				    const data = JSON.parse(e.data);
				    switch (data['frameType']) {
				        // 服务端ping客户端
				        case 1: // ping
				            this.websocketSend({type:"pong"});
				            break;
				        default:
				            this.$store.commit('catchSocketAction', data);
				            break;
				    }
				}
				
				
				this.websocket.onclose =event=> {
				    this.websocket.close();
				};
				Vue.prototype.$websocket = this.websocket;
            },
            websocketSend(agentData) {//数据发送
                var data=JSON.stringify(agentData);
                this.websocket.send(data);
            },
			
            playAudio () {
                const audio = document.getElementById('chatAudio');
                // 从头播放
                audio.currentTime = 0;
                audio.play();
            }
        },
        created() {			
			this.initWebSocket();
		}
    }
</script>
