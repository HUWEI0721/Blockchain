<template>
  <div class="navbar">
    <div class="navbar-container">
      <div class="logo-container">
        <img src="@/assets/logo/logo.png" class="logo">
        <span class="logo-title">野生水产品溯源系统</span>
      </div>
      
      <el-menu
        :default-active="activeMenu"
        class="horizontal-menu"
        mode="horizontal"
        background-color="#0d70c1"
        text-color="#fff"
        active-text-color="#ffd04b"
      >
        <div class="menu-wrapper">
          <sidebar-item
            v-for="route in routes"
            :key="route.path"
            :item="route"
            :base-path="route.path"
          />
        </div>
      </el-menu>

      <div class="right-menu">
        <el-dropdown class="avatar-container" trigger="click">
          <div class="avatar-wrapper">
            <img src="@/assets/logo/avatar.png" class="user-avatar">
            <i class="el-icon-caret-bottom" />
          </div>
          <el-dropdown-menu slot="dropdown" class="user-dropdown">
            <el-dropdown-item divided @click.native="logout">
              <span style="display:block;">登出</span>
            </el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import SidebarItem from './Sidebar/SidebarItem'
import variables from '@/styles/variables.scss'

export default {
  components: {
    SidebarItem
  },
  computed: {
    ...mapGetters([
      'avatar'
    ]),
    routes() {
      return this.$router.options.routes
    },
    activeMenu() {
      const route = this.$route
      const { meta, path } = route
      if (meta.activeMenu) {
        return meta.activeMenu
      }
      return path
    },
    variables() {
      return variables
    }
  },
  methods: {
    async logout() {
      await this.$store.dispatch('user/logout')
      this.$router.push(`/login?redirect=${this.$route.fullPath}`)
    }
  }
}
</script>

<style lang="scss" scoped>
.navbar {
  height: 60px;
  overflow: hidden;
  position: relative;
  background: #0d70c1;
  box-shadow: 0 1px 4px rgba(0,21,41,.08);

  .navbar-container {
    display: flex;
    align-items: center;
    justify-content: space-between;
    height: 100%;
    padding: 0 20px;
  }

  .logo-container {
    display: flex;
    align-items: center;
    margin-right: 30px;
    
    .logo {
      height: 40px;
      width: 40px;
      vertical-align: middle;
      border-radius: 50%;
      transition: transform 0.3s;
      margin-right: 10px;

      &:hover {
        transform: scale(1.1);
      }
    }

    .logo-title {
      color: #fff;
      font-size: 20px;
      font-weight: bold;
    }
  }

  .horizontal-menu {
    flex: 1;
    border: none;
    
    .menu-wrapper {
      display: flex;
    }

    :deep(.el-menu-item), :deep(.el-submenu__title) {
      height: 40px;
      line-height: 40px;
      color: #fff;
      padding: 0 20px;
      font-size: 16px;
      
      &:hover {
        background-color: #096dd9;
      }
      
      &.is-active {
        background-color: #096dd9;
        font-weight: bold;
      }
    }

    :deep(.el-submenu) {
      .el-menu {
        background-color: #0a60ad;
      }

      .el-menu-item {
        height: 50px;
        line-height: 50px;
        padding: 0 20px;
        min-width: 150px;
        
        &:hover {
          background-color: #0d70c1;
        }
      }
    }
  }

  .right-menu {
    margin-left: auto;
    
    .avatar-container {
      cursor: pointer;
      
      .avatar-wrapper {
        position: relative;

        .user-avatar {
          width: 40px;
          height: 40px;
          border-radius: 50%;
          transition: transform 0.3s;

          &:hover {
            transform: scale(1.1);
          }
        }

        .el-icon-caret-bottom {
          position: absolute;
          right: -15px;
          top: 20px;
          font-size: 12px;
          color: #fff;
          transition: transform 0.3s;
        }

        &:hover {
          .el-icon-caret-bottom {
            transform: rotate(180deg);
          }
        }
      }
    }
  }
}
</style>
