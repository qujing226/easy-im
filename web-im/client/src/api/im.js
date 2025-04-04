import request from '@/utils/request'

const imApi = {}

imApi.loginAPI= (params) =>{
	// console.log("请求信息");
	//
	// console.log(params);
  //   return request({
  //       url: 'v1/user/login',
  //       method: 'post',
  //       data: params
  //   })
	return request.get("http://"+window.location.host+"/data/login.json")

}

imApi.getSystemInfo= (params)  =>{
	return request.get("http://"+window.location.host+"/data/getSystemInfo.json")
}


imApi.getContactsAPI= (data)  =>{
	return request.get("http://"+window.location.host+"/data/getContacts.json")
}


imApi.getMessageListAPI= (data)  =>{
	return request.get("http://"+window.location.host+"/data/msg/"+data.toContactId+".json")
}


imApi.groupUserListAPI= (data)  =>{
	return request.get("http://"+window.location.host+"/data/groupuserlist.json")
}


imApi.getAllUserAPI= (data)  =>{
	return request.get("http://"+window.location.host+"/data/getAllUser.json")
}

imApi.getFileList= (data)  =>{
	return request.get("http://"+window.location.host+"/data/index.json")
}

export default imApi;
