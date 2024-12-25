<template>
  <div>
    <div class="chat-box">
      <lemon-imui
        :user="user"
        ref="IMUI"
        :width="curWidth"
        :height="curHeight"
        :theme="setting.theme"
        :hide-message-name="setting.hideMessageName"
        :hide-message-time="setting.hideMessageTime"
        :avatarCricle="setting.avatarCricle"
        @change-contact="handleChangeContact"
        @send="handleSend"
        :sendKey="setSendKey"
        :wrapKey="wrapKey"
        @pull-messages="handlePullMessages"
        style="min-height: 600px"
      >
        <!-- <lemon-imui :user="user" ref="IMUI" :width="curWidth" :height="curHeight" :contextmenu="contextmenu"
							:contact-contextmenu="contactContextmenu" :theme="setting.theme"
							:hide-message-name="setting.hideMessageName" :hide-message-time="setting.hideMessageTime"
							:avatarCricle="setting.avatarCricle" :sendKey="setSendKey" :wrapKey="wrapKey"
							@menu-avatar-click="openSetting" @change-contact="handleChangeContact"
							@pull-messages="handlePullMessages" @message-click="handleMessageClick" @send="handleSend"
							style="min-height:600px">	 -->
        <template #cover>
          <div @click="closeGroup">
            <div class="cover">
              <i class="lemon-icon-message"></i>
              <p><b>easy-im</b></p>
            </div>
          </div>
        </template>
        <template #sidebar-message="Contact">
          <span class="lemon-badge lemon-contact__avatar" @click="closeGroup">
            <span
              class="lemon-avatar"
              v-bind:class="{ 'lemon-avatar--circle': setting.avatarCricle }"
            >
              <img size="50px" :src="Contact.avatar"
            /></span>
            <span
              class="lemon-badge__label"
              v-if="Contact.unread > 0 && Contact.is_notice == 1"
              >{{ Contact.unread }}</span
            >
          </span>
          <div class="lemon-contact__inner" @click="closeGroup">
            <p class="lemon-contact__label">
              <span class="lemon-contact__name">
                <el-tag size="mini" v-if="Contact.is_group == 1">群聊</el-tag>
                {{ Contact.displayName }}
              </span>
              <span
                class="lemon-contact__time"
                v-text="formatTime(Contact.lastSendTime)"
              ></span>
            </p>
            <p class="lemon-contact__content lemon-last-content">
              <span class="lastContent">
                <span v-if="Contact.is_notice == 0 && Contact.unread > 0"
                  >[{{ Contact.unread }}条未读]</span
                >
                <span v-html="Contact.lastContent"></span>
              </span>
              <span
                class="el-icon-close-notification f-16"
                v-if="Contact.is_notice == 0"
              ></span>
            </p>
          </div>
        </template>
        <template #message-title="contact" style="color: red">
          <div class="message-title-box">
            <div>
              <span v-if="isEdit == false">
                <span class="displayName" v-if="is_group == 1">
                  <el-tag size="mini">群聊</el-tag> {{ contact.displayName
                  }}<span class="mr-5">({{ groupUserCount }})</span>
                  <el-tag
                    size="mini"
                    v-if="contact.setting.nospeak == 1"
                    type="warning"
                    >仅群管理员可发言</el-tag
                  >
                  <el-tag
                    size="mini"
                    v-if="contact.setting.nospeak == 2"
                    type="danger"
                    >全员禁言中</el-tag
                  >

                  <span class="mr-10" @click="openGroup">...</span>
                </span>
                <span class="displayName" v-if="is_group == 1">
                  {{ contact.displayName }}</span
                >
              </span>

              <input
                v-if="isEdit == true"
                v-model="displayName"
                class="editInput"
              />
            </div>
          </div>
        </template>
        <template #sidebar-message-fixedtop="instance">
          <div class="contact-fixedtop-box">
            <el-input
              placeholder="搜索"
              prefix-icon="el-icon-search"
              @blur="closeSearch"
              @focus="searchResult = true"
              v-model="keywords"
              class="input-with-select"
            >
            </el-input>
            <div
              style="margin-left: 10px"
              v-if="globalConfig.sysInfo.runMode == 2"
            >
              <el-dropdown @command="handleCommand">
                <el-button class="info"></el-button>
                <el-dropdown-menu slot="dropdown">
                  <el-dropdown-item
                    command="addGroup"
                    v-if="globalConfig.chatInfo.groupChat"
                    >创建群聊</el-dropdown-item
                  >
                </el-dropdown-menu>
              </el-dropdown>
            </div>
            <div
              style="margin-left: 10px"
              v-if="
                globalConfig.sysInfo.runMode == 1 &&
                globalConfig.chatInfo.groupChat
              "
            >
              <el-button
                title="创建群聊"
                icon="el-icon-plus"
                @click="openCreateGroup"
                circle
              ></el-button>
            </div>
            <div class="search-list" v-show="searchResult">
              <div
                v-for="(item, index) in searchList"
                :key="index"
                v-if="searchList.length > 0"
                class="search-list-item"
              >
                <lemon-contact
                  :contact="item"
                  @click="openChat(item.id, instance)"
                />
              </div>
              <div
                v-if="searchList.length == 0"
                style="margin: 20px"
                align="center"
              >
                暂无
              </div>
            </div>
          </div>
        </template>
        <template #sidebar-contact-fixedtop="instance">
          <div style="margin: 15px 10px">联系人</div>
        </template>
        <template #message-side="contact">
          <div
            class="slot-group-list"
            v-if="contact.is_group == 1"
            v-show="groupResult"
          >
            <div class="group-side-box">
              <div class="group-user">
                <div class="group-side-title">
                  <h4>群成员</h4>
                  <div>
                    <span
                      class="el-icon-circle-plus-outline f-18 cur-handle"
                      v-if="contact.role < 3 || contact.setting.invite == 1"
                      @click="openAddGroupUser"
                    ></span>
                  </div>
                </div>
                <hr />
                <div
                  class="group-user-body"
                  :style="'height:calc(' + curHeight + ' - 230px )'"
                  id="group-user"
                >
                  <el-scrollbar style="height: 100%">
                    <lemon-contact
                      class="user-list"
                      v-for="(item, index) in groupUser"
                      :key="index"
                      :contact="item"
                      v-lemon-contextmenu.contact="groupMenu"
                    >
                      <div class="user-avatar">
                        <el-avatar
                          :size="40"
                          :src="item.userInfo.avatar"
                        ></el-avatar>
                      </div>
                      <div class="user-name">
                        <span
                          v-if="item.userInfo.id == user.id"
                          class="fc-danger"
                          >{{ item.userInfo.displayName }}（我）</span
                        >
                        <span v-if="item.userInfo.id != user.id">{{
                          item.userInfo.displayName
                        }}</span>
                      </div>
                      <div class="user-role">
                        <i
                          class="el-icon-user-solid fc-danger"
                          title="创建者"
                          v-if="item.role == 1"
                        ></i>
                        <i
                          class="el-icon-user-solid fc-warning"
                          title="管理员"
                          v-if="item.role == 2"
                        ></i>
                      </div>
                    </lemon-contact>
                  </el-scrollbar>
                </div>
              </div>
            </div>
          </div>
        </template>
        <template #message-after="message">
          <span
            v-if="message.fromUser.id == user.id && message.is_group == 2"
            style="visibility: visible"
          >
            <span v-if="!message.is_read && message.status == 'succeed'">
              未读
            </span>
            <span
              v-if="message.is_read && message.status == 'succeed'"
              class="fc-success"
            >
              已读
            </span>
          </span>
        </template>
        <template #editor-footer>
          {{
            setting.sendKey == 1
              ? "使用 Enter 键发送消息"
              : "使用 Ctrl + Enter 键发送消息"
          }}
        </template>
      </lemon-imui>
    </div>
    <Group :visible.sync="createChatBox" :title="dialogTitle"></Group>
    <el-dialog
      title="发布公告"
      :visible.sync="noticeBox"
      :modal="true"
      width="500px"
      append-to-body
    >
      <el-input
        type="textarea"
        :rows="10"
        placeholder="请输入内容"
        v-model="notice"
      >
      </el-input>
      <span slot="footer" class="dialog-footer">
        <el-button>取 消</el-button>
        <el-button type="primary">确 定</el-button>
      </span>
    </el-dialog>
    <el-dialog
      title="消息管理器"
      :visible.sync="messageBox"
      :modal="true"
      width="800px"
      append-to-body
    >
    </el-dialog>
    <el-dialog
      title="群设置"
      :visible.sync="groupSetting"
      :modal="true"
      width="500px"
      append-to-body
    >
      <ChatSet :contact="contactSetting" :key="componentKey"></ChatSet>
    </el-dialog>
    <Socket ref="socket"></Socket>
  </div>
</template>

<script>
import { mapState } from "vuex";
import * as utils from "@/utils/index";
import Lockr from "lockr";
import Socket from "./socket";
import Group from "./group";
import { v4 } from "uuid";
const getTime = () => {
  return new Date().getTime();
};

const token = Lockr.get("authToken");
const id = Lockr.get("id");
export default {
  name: "app",
  components: {
    Socket,
    Group,
  },
  props: {
    width: {
      type: String,
      default: "70vw",
    },
    height: {
      type: String,
      default: "40vw",
    },
  },
  data() {
    var _this = this;
    return {
      curWidth: this.width,
      curHeight: this.height,
      unread: 0,
      curFile: null,
      componentKey: 1,
      searchResult: false,
      Result: false,
      createChatBox: false,
      forwardBox: false,
      noticeBox: false,
      messageBox: false,
      webrtcBox: false,
      groupSetting: false,
      VoiceStatus: false,
      contactSetting: {},
      groupUserCount: 0,
      dialogTitle: "创建群聊",
      isAdd: true,
      userIds: [],
      notice: "",
      searchList: [],
      keywords: "",
      displayName: "",
      oldName: "",
      isEdit: false,
      user: {
        id: id,
        displayName: "测试",
        avatar: "/static/img/avatar/62.jpg",
        account: token,
      },
      params: {
        page: 1,
        limit: 10,
      },
      is_group: 2,
      group_id: 0,
      contacts: [],
      allUser: [],
      groupUser: [],
      groupMenu: [],
      contactsToUid: new Map(),
    };
  },
  computed: {
    // 监听全局socket消息状态
    ...mapState({
      socketAction: (state) => state.socketAction,
      contactId: (state) => state.toContactId,
      contactSync: (state) => state.contactSync,
      setting: (state) => state.setting,
      userInfo: (state) => state.userInfo,
      globalConfig: (state) => state.globalConfig,
    }),
    formatTime() {
      return function (val) {
        return utils.timeFormat(val);
      };
    },
  },
  watch: {
    // 监听接收socket消息
    socketAction(val) {
      console.log("socketAction");
      console.log(val);

      const { IMUI } = this.$refs;
      const data = {
        id: val.data.MsgId,
        status: "succeed",
        type: "text",
        sendTime: val.data.sendTime,
        content: val.data.Content,
        toContactId: val.data.ConversationId,
        fromUser: {
          //如果 id == this.user.id消息会显示在右侧，否则在左侧
          id: val.formId,
          displayName: "00",
          avatar: this.contactsToUid.get(val.data.ConversationId).avatar,
        },
      };
      console.log(data);
      IMUI.appendMessage(data);
    },
  },
  created() {
    // 初始化用户
    let userInfo = this.$store.state.userInfo;
    if (userInfo) {
      this.user = {
        id: "1",
        displayName: "测试",
        avatar: "/static/img/avatar/62.jpg",
        account: token,
      };
    }
  },
  mounted() {
    if (this.searchResult) {
      document.addEventListener("click", function (e) {
        if (!that.$refs.configforms.contains(e.target)) {
          that.searchResult = false;
        }
      });
    }
    if (this.groupResult) {
      document.addEventListener("click", function (e) {
        if (!that.$refs.configforms.contains(e.target)) {
          that.groupResult = false;
        }
      });
    }
    // 初始化联系人
    this.getSimpleChat();
  },
  methods: {
    // 初始化聊天
    getSimpleChat(update) {
      this.$nextTick(() => {
        const IMUI = this.$refs.IMUI;
        this.IMUI = IMUI;
        // 获取联系人列表
        this.$api.imApi.getContactsAPI().then((res) => {
          const data = res.data;
          this.contacts = data;

          var contacts = [];
          var msg = {};
          // 重新渲染消息
          data.forEach((item, index) => {
            if (item.type) {
              msg.type = item.type;
              msg.content = item.lastContent;
              data[index]["lastContent"] = IMUI.lastContentRender(msg);
            }
            if (item.unread && !update) {
              this.unread += item.unread;
            }
            if (item.user_id != id) {
              contacts.push(item);
              this.contactsToUid.set(item.id, item);
            }
          });

          this.$store.commit("initContacts", contacts);
          // 设置置顶人
          this.getChatTop(contacts);
          IMUI.initContacts(contacts);
          // 初始化左侧菜单栏
          this.initMenus(IMUI);
        });
      });
    },
    // 初始化菜单
    initMenus(IMUI) {
      let menus = [
        {
          name: "messages",
          unread: this.unread,
        },
        {
          name: "contacts",
        },
      ];
      IMUI.initMenus(menus);
    },
    // 获取置顶联系人列表
    getChatTop() {
      var list = this.contacts;
      var topList = [];
      for (var i = 0; i < list.length; i++) {
        if (list[i].is_top == 1) {
          topList.push(list[i]);
        }
      }
      this.chatTopList = topList;
    },
    // 获取联系人
    getContact(id) {
      const { IMUI } = this.$refs;
      const contactList = IMUI.getContacts();
      for (var i = 0; i < contactList.length; i++) {
        if (contactList[i].id == id) {
          return contactList[i];
        }
      }
    },
    // 打开聊天窗口
    openChat(contactId, instance) {
      this.keywords = "";
      instance.changeContact(contactId);
    },
    // 切换聊天窗口时要检测那些消息未读
    handleChangeContact(contact, instance) {
      instance.updateContact({
        id: contact.id,
        unread: 0,
      });
      // 将未读的总数减去当前选择的聊天
      this.unread -= contact.unread;
      const { IMUI } = this.$refs;
      this.initMenus(IMUI);
      // 聊天记录列表恢复到最初第一页
      this.params.page = 1;
      this.displayName = contact.displayName;
      this.oldName = contact.displayName;
      this.currentChat = contact;
      // 如果是群聊，拉取群成员列表，如果刚才拉取过，现在就不用拉取了
      if (contact.is_group == 1 && this.group_id != contact.id) {
        this.getGroupUserList(contact.id);
      }
      //切换聊天后全局设置是否是群聊或者单聊
      this.is_group = contact.is_group;
      // 如果是团队id，保存当前团队id避免下次切换回来的时候重复请求成员列表
      if (this.is_group == 1) {
        this.group_id = contact.id;
        this.notice = contact.notice;
      }
      var data = [];

      // 获取当前聊天窗口的所有消息
      var messages = IMUI.getMessages(contact.id);
      for (var i = 0; messages.length > i; i++) {
        if (
          messages[i].is_read == 0 &&
          messages[i].fromUser.id != this.user.id
        ) {
          data.push(messages[i]);
        }
      }
      // 如果有未读的消息，需要将消息修改为已读
      if (data.length > 0) {
        this.$api.imApi
          .setMsgIsReadAPI({
            is_group: contact.is_group,
            toContactId: contact.id,
            messages: data,
            fromUser: contact.id,
          })
          .then((res) => {
            if (res.code == 0) {
              this.setLocalMsgIsRead(data);
            }
          });
      }
      instance.closeDrawer();
    },
    // 发送聊天消息
    handleSend(message, next, file) {
      console.log(message);
      const data = {
        id: message.id,
        method: "conversation.chat",
        data: {
          chatType: this.is_group,
          recvId: this.contactsToUid.get(message.toContactId).user_id,
          msg: {
            content: message.content,
          },
        },
      };
      console.log(data);

      this.$refs.socket.websocketSend(data);

      next();
    },
    // 拉取聊天记录
    handlePullMessages(contact, next, instance) {
      let params = this.params;
      params.toContactId = contact.id;
      params.is_group = contact.is_group;
      this.$api.imApi
        .getMessageListAPI(params)
        .then((res) => {
          this.params.page++;
          let isEnd = false;
          let messages = res.data;
          if (messages.length < this.params.limit) {
            isEnd = true;
          }
          next(messages, isEnd);
        })
        .catch((error) => {
          next([], true);
        });
      return true;
    },
    // 打开创建团队的窗口
    openCreateGroup() {
      this.isAdd = true;
      this.dialogTitle = "创建群聊";
      this.createChatBox = true;
    },

    // 打开添加群成员的窗口
    openAddGroupUser() {
      var user_ids = utils.arrayToString(this.groupUser, "user_id");
      this.isAdd = false;
      this.userIds = user_ids;
      this.dialogTitle = "添加群成员";
      this.createChatBox = true;
    },
    // 封装循环请求
    fn(formData) {
      return new Promise((resolve, reject) => {
        this.$api.imApi
          .sendMessageAPI(formData)
          .then((res) => {
            if (res.code === 0) {
              resolve(res);
            } else {
              this.$message.error(res.msg);
            }
          })
          .catch((err) => {});
      });
    },
    async test(formData) {
      let n = await this.fn(formData);
      return n;
    },

    // 获取群聊成员列表
    getGroupUserList(group_id) {
      this.$api.imApi
        .groupUserListAPI({
          group_id: group_id,
        })
        .then((res) => {
          if (res.code == 0) {
            var data = res.data;
            this.groupUser = data;
            this.groupUserCount = data.length;
          }
        });
    },
    // 关闭搜索结果
    closeSearch() {
      var that = this;
      setTimeout(function () {
        that.searchResult = false;
      }, 300);
    },
    openGroup() {
      if (this.groupResult) {
        this.groupResult = false;
      } else {
        this.groupResult = true;
      }
    },
    closeGroup() {
      var that = this;
      setTimeout(function () {
        that.groupResult = false;
      }, 300);
    },
    // 将本地消息设置为已读
    setLocalMsgIsRead(message) {
      const { IMUI } = this.$refs;
      for (let i = 0; message.length > i; i++) {
        const data = {
          id: message[i]["id"],
          is_read: 1,
          status: "succeed",
          content: message[i]["content"] + " ",
        };
        IMUI.updateMessage(data);
      }
    },
    openMessageBox() {
      this.messageBox = true;
      // 组件重置
      this.componentKey += 1;
    },
    // 打开设置中心
    openSetting() {
      const { IMUI } = this.$refs;
      IMUI.changeMenu("setting");
    },
    handleCommand(e) {
      if (e == "addGroup") {
        this.openCreateGroup();
      } else {
        this.addFriendBox = true;
      }
    },
    wrapKey(e) {
      return this.setting.sendKey == 1
        ? e.keyCode == 13 && e.ctrlKey
        : e.keyCode == 13 && !e.ctrlKey && !e.shiftKey;
    },
    // 设置发送键
    setSendKey(e) {
      return this.setting.sendKey == 1
        ? e.keyCode == 13 && !e.ctrlKey && !e.shiftKey
        : e.keyCode == 13 && e.ctrlKey;
    },
    //自定义消息的发送
    diySendMessage(message, file) {
      const { IMUI } = this.$refs;
      IMUI.appendMessage(message, true);
      this.handleSend(
        message,
        function () {
          var replaceMessage =
            arguments.length > 0 && arguments[0] !== undefined
              ? arguments[0]
              : {
                  status: "succeed",
                };
          IMUI.updateContact({
            id: message.toContactId,
            lastContent: IMUI.lastContentRender(message),
            lastSendTime: message.sendTime,
          });
          IMUI.CacheDraft.remove(message.toContactId);
          IMUI.updateMessage(Object.assign(message, replaceMessage));
        },
        file
      );
    },
    // 退出聊天室
    logout() {
      this.$confirm("你确定要退出聊天室吗?", "提示", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      })
        .then(() => {
          this.$store.dispatch("LogOut").then(() => {
            this.$router.push({
              path: "/login",
            });
          });
        })
        .catch(() => {
          this.$message({
            type: "info",
            message: "已取消退出",
          });
        });
    },
  },
};
</script>
<style scoped lang="scss">
.mr-10 {
  margin-right: 0px;
  font-size: 33px;
  right: 23px;
  top: -3px;
  position: absolute;
}

.main-container {
  display: flex;
  -webkit-box-pack: center;
  -ms-flex-pack: center;
  justify-content: center;
  -webkit-box-align: center;
  -ms-flex-align: center;
  align-items: center;
  width: 100%;
  height: 300vh;
  background-size: cover;
}

.messageBoxStyle {
  position: fixed;
  top: 0;
  left: 0;
  height: 250vh;
  width: 100%;
  z-index: 1999;
  background: rgba(0, 0, 0, 0.5);
  overflow-y: visible;

  .el-dialog__wrapper {
    display: flex;
    align-items: center;
  }
}

.chat-box {
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.12), 0 0 6px rgba(0, 0, 0, 0.04);
}

.cover {
  text-align: center;
  user-select: none;
  position: absolute;
  top: 45%;
  left: 50%;
  transform: translate(-50%, -50%);

  i {
    font-size: 84px;
    color: #e6e6e6;
  }

  p {
    font-size: 18px;
    color: #ddd;
    line-height: 50px;
  }
}

.displayName {
  font-size: 16px;
  visibility: visible;
}

.contact-fixedtop-box {
  margin: 15px 10px;
  display: flex;
  align-items: center;
  justify-content: space-between;
  visibility: visible;
  position: relative;
}

.search-list {
  background: #fff;
  position: absolute;
  z-index: 99;
  top: 40px;
  width: 230px;
  height: 300px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.12);
  overflow: auto;
  border: solid 1px #e6e6e6;

  .search-list-item :hover {
    background: #f1f1f1;
  }

  .lemon-contact {
    background: #fff;
  }
}

.chat-top-list {
  display: flex;
  padding: 0 5px 10px 10px;
  justify-content: flex-start;
  flex-wrap: wrap;
}

.message-title-box {
  display: flex;
  align-items: center;
  justify-content: space-between;
  visibility: visible;
}

.message-title-tools {
  font-size: 30px;
  color: #999999;
  letter-spacing: 5px;
  cursor: pointer;
}

.editInput {
  font-size: 18px;
  border: none;
  width: 400px;
}

.editInput:focus {
  outline: -webkit-focus-ring-color auto 0px;
}

.lemon-last-content {
  display: flex;
  justify-content: space-between;

  .lastContent {
    width: 200px !important;
    text-overflow: ellipsis;
    overflow: hidden;
    white-space: nowrap;
  }
}

.slot-group-list {
  background: #fff;
  width: 200px;
  border-left: solid 1px #e6e6e6;
  height: 100%;
  white-space: initial;

  .group-side-box {
    .group-side-title {
      padding: 5px 10px;
    }

    .group-side-body {
      height: 150px;
      padding: 10px;
      cursor: pointer;
      overflow: hidden;
      text-overflow: ellipsis;
      display: -webkit-box;
      -webkit-line-clamp: 5;
      -webkit-box-orient: vertical;
    }

    .group-user-body {
      min-height: 410px;

      .user-list {
        display: flex;
        flex-direction: row;
        align-items: center;
        flex-wrap: nowrap;
        justify-content: flex-start;
        padding: 5px;
        background: #fff;

        .user-avatar {
          margin: 3px 8px 0 0;
          line-height: 10px;
        }

        .user-name {
          width: 110px;
        }

        .user-role {
          width: 20px;
        }
      }

      .user-list:hover {
        background: #e6e6e6;
      }
    }
  }
}

.group-side-title {
  display: flex;
  flex-direction: row;
  align-items: center;
  flex-wrap: nowrap;
  justify-content: space-between;
}

.group-notice {
  height: 140px;
}

.group-user {
  min-height: 300px;
  overflow: auto;
}

.lemon-avatar {
  width: 40px;
  height: 50px;
  line-height: 40px;
  font-size: 30px;
}

.lemon-menu__item {
  color: #ffffff;
  cursor: pointer;
  padding: 14px 10px;
  max-width: 100%;
  word-break: break-all;
  word-wrap: break-word;
  white-space: pre-wrap;
}

.lemon-menu {
  -webkit-box-align: center;
  -ms-flex-align: center;
  align-items: center;
  width: 60px;
  background: #409eff;
  padding: 15px 0;
  position: relative;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}
</style>
