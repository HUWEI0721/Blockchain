<template>
  <div class="login-container">
    <div class="login-box">
      <div class="login-content">
        <div class="login-logo">
          <img src="@/assets/logo/logo.png" alt="logo">
          <span class="logo-text">野生水产品溯源系统</span>
        </div>

        <el-form ref="loginForm" :model="loginForm" :rules="loginRules" class="login-form" auto-complete="on" label-position="left">
          <div v-show="isLoginPage" class="form-container">

            <div class="form-body">
              <el-form-item prop="username">
                <span class="svg-container">
                  <svg-icon icon-class="user" />
                </span>
                <el-input
                  ref="username"
                  v-model="loginForm.username"
                  placeholder="请输入账号"
                  name="username"
                  type="text"
                  tabindex="1"
                  auto-complete="on"
                />
              </el-form-item>

              <el-form-item prop="password" :class="{'password-item': passwordType === 'password'}">
                <span class="svg-container">
                  <svg-icon icon-class="password" />
                </span>
                <el-input
                  :key="passwordType"
                  ref="password"
                  v-model="loginForm.password"
                  :type="passwordType"
                  placeholder="请输入密码"
                  name="password"
                  tabindex="2"
                  auto-complete="on"
                  @keyup.enter.native="handleLogin"
                />
                <span class="show-pwd" @click="showPwd">
                  <svg-icon :icon-class="passwordType === 'password' ? 'eye' : 'eye-open'" />
                </span>
              </el-form-item>

              <div class="button-container">
                <el-button :loading="loading" type="primary" class="login-button" @click.native.prevent="handleLogin">登 录</el-button>
                <el-button class="register-button" @click="handleRegister">注 册</el-button>
              </div>
            </div>
          </div>

          <div v-show="!isLoginPage" class="form-container">
            <h2 class="login-title">用户注册</h2>

            <el-form-item prop="username">
              <span class="svg-container">
                <svg-icon icon-class="user" />
              </span>
              <el-input
                v-model="registerForm.username"
                placeholder="请输入账号"
                name="username"
                type="text"
                auto-complete="on"
              />
            </el-form-item>

            <el-form-item prop="password">
              <span class="svg-container">
                <svg-icon icon-class="password" />
              </span>
              <el-input
                :key="passwordType"
                ref="password"
                v-model="registerForm.password"
                :type="passwordType"
                placeholder="请输入密码"
                name="password"
                auto-complete="on"
              />
              <span class="show-pwd" @click="showPwd">
                <svg-icon :icon-class="passwordType === 'password' ? 'eye' : 'eye-open'" />
              </span>
            </el-form-item>

            <el-form-item prop="password2">
              <span class="svg-container">
                <svg-icon icon-class="password" />
              </span>
              <el-input
                v-model="registerForm.password2"
                placeholder="请再次输入密码"
                name="password"
                auto-complete="on"
                :type="passwordType"
              />
              <span class="show-pwd" @click="showPwd">
                <svg-icon :icon-class="passwordType === 'password' ? 'eye' : 'eye-open'" />
              </span>
            </el-form-item>

            <el-form-item>
              <el-select v-model="registerForm.userType" placeholder="请选择角色" class="role-select">
                <el-option
                  v-for="item in options"
                  :key="item.value"
                  :label="item.label"
                  :value="item.value"
                />
              </el-select>
            </el-form-item>

            <div class="button-container">
              <el-button class="back-button" @click="handleRegister">返 回</el-button>
              <el-button :loading="loading" type="primary" class="register-submit" @click.native.prevent="submitRegister">提交注册</el-button>
            </div>
          </div>
        </el-form>
      </div>

      <div class="login-footer">
        <span>Copyright © 2024 野生水产品溯源系统. All rights reserved.</span>
      </div>
    </div>
  </div>
</template>

<script>

export default {
  name: 'Login',
  data() {
    return {
      loginForm: {
        username: '',
        password: ''
      },
      loginRules: {
        username: [{ required: true }],
        password: [{ required: true }]
      },
      loading: false,
      passwordType: 'password',
      redirect: undefined,
      isLoginPage: true,
      registerForm: {
        username: '',
        password: '',
        password2: '',
        userType: ''
      },
      options: [{
        value: '渔民',
        label: '渔民'
      }, {
        value: '加工厂',
        label: '加工厂'
      }, {
        value: '物流司机信息',
        label: '物流司机信息'
      }, {
        value: 'keyword_29',
        label: '销售终端'
      }, {
        value: '消费者',
        label: '消费者'
      }]
    }
  },
  watch: {
    $route: {
      handler: function(route) {
        this.redirect = route.query && route.query.redirect
      },
      immediate: true
    }
  },
  methods: {
    showPwd() {
      if (this.passwordType === 'password') {
        this.passwordType = ''
      } else {
        this.passwordType = 'password'
      }
      this.$nextTick(() => {
        this.$refs.password.focus()
      })
    },
    handleLogin() {
      this.loading = true
      this.$store.dispatch('user/login', this.loginForm).then(() => {
        this.$router.push({ path: this.redirect || '/' })
        this.loading = false
      }).catch(() => {
        this.loading = false
      })
    },
    handleRegister() {
      // 切换登录注册
      this.isLoginPage = !this.isLoginPage
    },
    submitRegister() {
      if (this.registerForm.password !== this.registerForm.password2) {
        this.$message.error('两次密码不一致')
        return
      }
      const loading = this.$loading({
        lock: true,
        text: '注册中...',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      this.$store.dispatch('user/register', this.registerForm).then(response => {
        this.$router.push({ path: this.redirect || '/' })
        this.loading = false
        this.$message({
          message: '注册成功，链上交易ID：' + response.txid,
          type: 'success'
        })
        loading.close()
        this.handleRegister()
      }).catch(() => {
        loading.close()
      })
    }
  }
}
</script>

<style lang="scss">
.el-form-item {
  display: flex;
  align-items: center;
}

.svg-container {
  padding: 6px 5px 6px 15px;
  color: #889aa4;
  vertical-align: middle;
  width: 30px;
  display: inline-block;
  margin-right: 10px;
}

.el-form-item__content {
  flex: 1;
  display: flex;
  align-items: center;
}

.el-input {
  &__inner {
    background: rgba(255,255,255,0.8);
    border-color: rgba(255,255,255,0.5);
    color: #333;

    &::placeholder {
      color: #999;
    }
  }
}

.el-form-item {
  &.password-item {
    .el-input__inner {
      padding-right: 40px;
    }
  }
}

.show-pwd {
  position: absolute;
  right: 10px;
  top: 7px;
  font-size: 16px;
  color: #889aa4;
  cursor: pointer;
  user-select: none;
}
</style>

<style lang="scss" scoped>
.login-container {
  min-height: 100%;
  width: 100%;
  background-image: linear-gradient(rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)), url("../../assets/login_images/nature.webp");
  background-size: cover;
  background-position: center;
  display: flex;
  justify-content: center;
  align-items: center;
  position: relative;
  overflow: hidden;

  &::before {
    content: '';
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    backdrop-filter: blur(2px);
  }

  .login-box {
    position: relative;
    width: 520px;
    background: transparent;

    .login-content {
      position: relative;
      padding: 40px;
      background: rgba(255, 255, 255, 0.1);
      border-radius: 30px;
      box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.37);
      backdrop-filter: blur(8px);
      border: 1px solid rgba(255, 255, 255, 0.18);
      overflow: hidden;

      &::before {
        content: '';
        position: absolute;
        top: -50%;
        left: -50%;
        width: 200%;
        height: 200%;
        background: radial-gradient(circle at center, rgba(255,255,255,0.1) 0%, transparent 60%);
        transform: rotate(45deg);
        pointer-events: none;
      }

      .login-logo {
        display: flex;
        align-items: center;
        justify-content: center;
        margin-bottom: 40px;

        img {
          height: 80px;
          margin-right: 20px;
          filter: drop-shadow(0 0 10px rgba(255,255,255,0.3));
          animation: float 6s ease-in-out infinite;
        }

        .logo-text {
          font-size: 28px;
          font-weight: bold;
          color: #fff;
          text-shadow: 0 0 10px rgba(13,112,193,0.5);
        }
      }

      .login-form {
        .form-container {
          animation: fadeIn 0.5s ease;
        }

        .form-header {
          text-align: center;
          margin-bottom: 40px;

          .login-title {
            color: #fff;
            font-size: 24px;
            font-weight: 500;
            margin-bottom: 10px;
            text-shadow: 0 0 10px rgba(13,112,193,0.5);
          }

          .login-subtitle {
            color: rgba(255,255,255,0.7);
            font-size: 14px;
          }
        }

        .form-body {
          background: rgba(255, 255, 255, 0.05);
          padding: 30px;
          border-radius: 20px;
          backdrop-filter: blur(4px);

          .el-form-item {
            margin-bottom: 25px;
            border-radius: 10px;
            overflow: hidden;
            transition: all 0.3s ease;

            &:hover {
              transform: translateY(-2px);
              box-shadow: 0 5px 15px rgba(0,0,0,0.2);
            }
          }

          .button-container {
            display: flex;
            justify-content: space-between;
            margin-top: 30px;

            .login-button, .register-button, .back-button, .register-submit {
              width: 48%;
              height: 44px;
              font-size: 16px;
              letter-spacing: 4px;
              border-radius: 10px;
              transition: all 0.3s ease;
            }

            .login-button, .register-submit {
              background: #0d70c1;
              border-color: #0d70c1;

              &:hover {
                background: #096dd9;
                border-color: #096dd9;
                transform: translateY(-2px);
                box-shadow: 0 5px 15px rgba(13,112,193,0.3);
              }
            }

            .register-button, .back-button {
              background: rgba(255,255,255,0.2);
              border-color: rgba(255,255,255,0.3);
              color: #fff;

              &:hover {
                background: rgba(255,255,255,0.3);
                border-color: rgba(255,255,255,0.4);
                transform: translateY(-2px);
                box-shadow: 0 5px 15px rgba(255,255,255,0.2);
              }
            }
          }
        }
      }
    }
  }

  .login-footer {
    position: absolute;
    bottom: 20px;
    width: 100%;
    text-align: center;
    color: rgba(255,255,255,0.7);
    font-size: 14px;
    text-shadow: 0 0 10px rgba(0,0,0,0.3);
  }
}

@keyframes float {
  0% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(-10px);
  }
  100% {
    transform: translateY(0px);
  }
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.role-select {
  width: 100%;

  .el-input__inner {
    background: rgba(255,255,255,0.1);
    border-color: rgba(255,255,255,0.2);
    color: #fff;

    &::placeholder {
      color: rgba(255,255,255,0.5);
    }
  }

  .el-select__caret {
    color: rgba(255,255,255,0.5);
  }
}
</style>
