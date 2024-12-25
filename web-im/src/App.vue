<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
import lockr from 'lockr';

export default {
  name: 'App',
  data () {
    return {
      msg: 'Welcome to Your Vue.js App'
    }
  },
  created() {
    let userInfo=lockr.get('UserInfo');
    let Config=lockr.get('globalConfig');
    if(Config){
      this.$store.commit('setGlobalConfig', Config);
    }
	this.$store.dispatch('getSystemInfo').then(res=>{});
  }
}
</script>

<style lang="scss">
@import "assets/scss/style";

  .lemon-container .lemon-container__title{
    border-bottom:solid 1px #e6e6e6 !important
  }
  .el-transfer-panel .el-transfer-panel__body {
    height: 280px !important;
  }
  .lemon-editor__submit .lemon-button{
    background: var(--el-color-primary);
    color:#fff;
  }
   .lemon-editor__submit .lemon-button:hover{
    background: var(--el-color-primary);
    color:#fff !important;
    border:solid 1px var(--el-color-primary) !important;
  }
  .lemon-editor__submit button[disabled],.lemon-editor__submit button[disabled]:hover{
    background: #fff;
    color:#aaa;
    border:solid 1px #aaa;
  }
  .lemon-wrapper .lemon-menu{
	  -webkit-box-align: center;
	  -ms-flex-align: center;
	  align-items: center;
	  width: 60px;
	  background: #66b1ff;
	  padding: 15px 0;
	  position: relative;
	  -webkit-user-select: none;
	  -moz-user-select: none;
	  -ms-user-select: none;
	  user-select: none;
  }
  .lemon-wrapper .lemon-menu .lemon-menu__item {
      color: #ffffff;
      cursor: pointer;
      padding: 14px 10px;
      max-width: 100%;
      word-break: break-all;
      word-wrap: break-word;
      white-space: pre-wrap;
  }
  body,html,#app{
    height: 100%;
    width:100%
  }

  .el-scrollbar{
      height: 100% !important;
  }

  .el-scrollbar__wrap {
      overflow-x: hidden !important;
  }

  .el-container{
    overflow: auto;
  }

  .el-select-dropdown__wrap {
    margin-bottom: 0 !important;
  }

  .lemon-contact{
    padding:10px;
  }


  hr {
    height: 1px;
    background-color: #e6e6e6;
    border: none;
  }


</style>
