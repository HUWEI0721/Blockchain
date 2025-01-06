<!-- Start of Selection -->
<template>
  <div class="uplink-container" :class="roleClass">
    <el-card class="user-info-card" shadow="hover">
      <div class="user-info">
        <div class="info-header">
          <h2>用户信息</h2>
          <div class="role-icon">
            <i :class="roleIcon"></i>
          </div>
        </div>
        <div class="info-content">
          <div class="info-item">
            <i class="el-icon-user"></i>
            <span class="user-text">当前用户: <strong>{{ name }}</strong></span>
          </div>
          <el-divider direction="vertical"></el-divider>
          <div class="info-item">
            <i class="el-icon-s-custom"></i>
            <span class="role-text">用户角色: <strong>{{ userType }}</strong></span>
          </div>
        </div>
      </div>
    </el-card>

    <el-card class="main-form-card" shadow="hover">
      <el-form ref="form" :model="tracedata" label-width="120px" size="small">

        <el-form-item v-show="userType!='渔民'&userType!='消费者'" label="溯源码:" class="trace-code-item">
          <el-input v-model="tracedata.traceability_code" placeholder="请输入溯源码">
            <template slot="prepend"><i class="el-icon-key"></i></template>
          </el-input>
        </el-form-item>

        <!-- 渔民信息录入区域 -->
        <div v-show="userType=='渔民'" class="role-form-section">
          <h3><i class="el-icon-ship"></i> 渔民信息录入 <i class="el-icon-fishing role-specific-icon"></i></h3>
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="水产品名称:">
                <el-input v-model="tracedata.Fisherman_input.Sf_seafoodName" placeholder="请输入水产品名称"/>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="捕捞区域:">
                <el-input v-model="tracedata.Fisherman_input.Sf_origin" placeholder="请输入捕捞区域"/>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="捕捞起始时间:">
                <el-date-picker v-model="tracedata.Fisherman_input.Sf_salvageTime" type="datetime" placeholder="选择日期时间" style="width: 100%"/>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="出水时间:">
                <el-date-picker v-model="tracedata.Fisherman_input.Sf_outOfWaterTime" type="datetime" placeholder="选择日期时间" style="width: 100%"/>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="渔民名称:">
                <el-input v-model="tracedata.Fisherman_input.Sf_fishermanName" placeholder="请输入渔民名称"/>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="捕捞时长:">
                <el-input :value="calculateDuration" readonly>
                  <template slot="append">小时</template>
                </el-input>
              </el-form-item>
            </el-col>
          </el-row>
        </div>

        <!-- 加工厂信息录入区域 -->
        <div v-show="userType=='加工厂'" class="role-form-section">
          <h3><i class="el-icon-office-building"></i> 加工厂信息录入 <i class="el-icon-box role-specific-icon"></i></h3>
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="商品名称:">
                <el-input v-model="tracedata.Factory_input.Fac_productName" placeholder="请输入商品名称"/>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="加工批次:">
                <el-input v-model="tracedata.Factory_input.Fac_productionbatch" placeholder="请输入加工批次"/>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="加工时间:">
                <el-date-picker v-model="tracedata.Factory_input.Fac_productionTime" type="datetime" placeholder="选择日期时间" style="width: 100%"/>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="加工厂名称:">
                <el-input v-model="tracedata.Factory_input.Fac_factoryName" placeholder="请输入加工厂名称"/>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="联系电话:">
                <el-input v-model="tracedata.Factory_input.Fac_contactNumber" placeholder="请输入联系电话"/>
              </el-form-item>
            </el-col>
          </el-row>
        </div>

        <!-- 物流司机信息录入区域 -->
        <div v-show="userType=='物流司机信息'" class="role-form-section">
          <h3><i class="el-icon-truck"></i> 物流信息录入 <i class="el-icon-van role-specific-icon"></i></h3>
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="司机姓名:">
                <el-input v-model="tracedata.Driver_input.Dr_name" placeholder="请输入司机姓名"/>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="冷链温度:">
                <el-input v-model="tracedata.Driver_input.Dr_age" placeholder="请输入冷链温度">
                  <template slot="append">°C</template>
                </el-input>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="联系电话:">
                <el-input v-model="tracedata.Driver_input.Dr_phone" placeholder="请输入联系电话"/>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="车牌号:">
                <el-input v-model="tracedata.Driver_input.Dr_carNumber" placeholder="请输入车牌号"/>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="运输工具:">
                <el-input v-model="tracedata.Driver_input.Dr_transport" placeholder="请输入运输工具"/>
              </el-form-item>
            </el-col>
          </el-row>
        </div>

        <!-- 销售点信息录入区域 -->
        <div v-show="userType=='keyword_29'" class="role-form-section">
          <h3><i class="el-icon-shopping-cart-full"></i> 销售点信息录入 <i class="el-icon-goods role-specific-icon"></i></h3>
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="存入时间:">
                <el-date-picker v-model="tracedata.Shop_input.Sh_storeTime" type="datetime" placeholder="选择日期时间" style="width: 100%"/>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="销售时间:">
                <el-date-picker v-model="tracedata.Shop_input.Sh_sellTime" type="datetime" placeholder="选择日期时间" style="width: 100%"/>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row :gutter="20">
            <el-col :span="12">
              <el-form-item label="销售点名称:">
                <el-input v-model="tracedata.Shop_input.Sh_shopName" placeholder="请输入销售点名称"/>
              </el-form-item>
            </el-col>
            <el-col :span="12">
              <el-form-item label="销售点位置:">
                <el-input v-model="tracedata.Shop_input.Sh_shopAddress" placeholder="请输入销售点位置"/>
              </el-form-item>
            </el-col>
          </el-row>
          <el-row>
            <el-col :span="12">
              <el-form-item label="销售点电话:">
                <el-input v-model="tracedata.Shop_input.Sh_shopPhone" placeholder="请输入销售点电话"/>
              </el-form-item>
            </el-col>
          </el-row>
        </div>

        <!-- 图片上传区域 -->
        <div v-show="userType != '消费者'" class="upload-section">
          <h3><i class="el-icon-picture"></i> 图片上传</h3>
          <el-upload
            class="upload-area"
            action="#"
            :on-preview="handlePreview"
            :on-remove="handleRemove"
            :before-remove="beforeRemove"
            :on-change="handleChange"
            :file-list="fileList"
            :auto-upload="false"
            :limit="1"
            list-type="picture-card">
            <template v-if="fileList.length < 1">
              <i class="el-icon-plus upload-icon"></i>
            </template>
          </el-upload>
          <div class="upload-tip" v-if="fileList.length < 1">
            <i class="el-icon-info-circle"></i>
            <span>图片大小不超过 500KB，建议尺寸 1080x1080px</span>
          </div>
        </div>

        <div class="form-footer">
          <el-alert
            v-show="userType == '消费者'"
            title="消费者没有权限录入！请使用溯源功能!"
            type="warning"
            :closable="false">
          </el-alert>
          <el-button
            v-show="userType != '消费者'"
            type="success"
            icon="el-icon-upload2"
            @click="submittracedata">
            提交数据上链
          </el-button>
        </div>
      </el-form>
    </el-card>

    <!-- 图片预览对话框 -->
    <el-dialog :visible.sync="previewVisible">
      <img width="100%" :src="previewUrl" alt="Preview">
    </el-dialog>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { uplink } from '@/api/trace'
import Compressor from 'compressorjs';

export default {
  name: 'Uplink',
  data() {
    return {
      tracedata: {
        traceability_code: '',
        Fisherman_input: {
          Sf_seafoodName: '',
          Sf_origin: '',
          Sf_salvageTime: '',
          Sf_outOfWaterTime: '',
          Sf_fishermanName: ''
        },
        Factory_input: {
          Fac_productName: '',
          Fac_productionbatch: '',
          Fac_productionTime: '',
          Fac_factoryName: '',
          Fac_contactNumber: ''
        },
        Driver_input: {
          Dr_name: '',
          Dr_age: '',
          Dr_phone: '',
          Dr_carNumber: '',
          Dr_transport: ''
        },
        Shop_input: {
          Sh_storeTime: '',
          Sh_sellTime: '',
          Sh_shopName: '',
          Sh_shopAddress: '',
          Sh_shopPhone: ''
        }
      },
      fileList: [],
      loading: false,
      previewVisible: false,
      previewUrl: ''
    }
  },
  computed: {
    ...mapGetters([
      'name',
      'userType'
    ]),
    calculateDuration() {
      if (this.tracedata.Fisherman_input.Sf_salvageTime && this.tracedata.Fisherman_input.Sf_outOfWaterTime) {
        const startTime = new Date(this.tracedata.Fisherman_input.Sf_salvageTime).getTime()
        const endTime = new Date(this.tracedata.Fisherman_input.Sf_outOfWaterTime).getTime()
        const diffHours = Math.round((endTime - startTime) / (1000 * 60 * 60))
        return diffHours >= 0 ? diffHours : 0
      }
      return 0
    },
    roleClass() {
      switch(this.userType) {
        case '渔民':
          return 'fisherman-bg'
        case '加工厂':
          return 'factory-bg'
        case '物流司机信息':
          return 'driver-bg'
        case 'keyword_29':
          return 'shop-bg'
        default:
          return 'default-bg'
      }
    },
    roleIcon() {
      switch(this.userType) {
        case '渔民':
          return 'el-icon-ship'
        case '加工厂':
          return 'el-icon-office-building'
        case '物流司机信息':
          return 'el-icon-truck'
        case 'keyword_29':
          return 'el-icon-shopping-cart-full'
        default:
          return 'el-icon-user'
      }
    }
  },
  methods: {
    handleRemove(file, fileList) {
      console.log(file, fileList);
    },
    handlePreview(file) {
      this.previewUrl = file.url;
      this.previewVisible = true;
    },
    handleChange(file, fileList) {
      new Compressor(file.raw, {
        quality: 0.2, // 压缩质量，0到1之间
        convertSize: Infinity, // 确保所有图片都被转换
        mimeType: 'image/jpeg', // 转换为jpg格式

        success: (compressedFile) => {
          this.fileList = [{ raw: compressedFile, url: URL.createObjectURL(compressedFile) }];
        },
        error(err) {
          console.error(err.message);
        },
      });
    },
    beforeRemove(file, fileList) {
      return this.$confirm(`确定移除 ${file.name}？`);
    },
    submittracedata() {
      console.log(this.tracedata)
      const loading = this.$loading({
        lock: true,
        text: '数据上链中...',
        spinner: 'el-icon-loading',
        background: 'rgba(0, 0, 0, 0.7)'
      })
      var formData = new FormData()
      formData.append('traceability_code', this.tracedata.traceability_code)
      // 根据不同的用户给arg1、arg2、arg3..赋值,
      switch (this.userType) {
        case '渔民':
          formData.append('arg1', this.tracedata.Fisherman_input.Sf_seafoodName)
          formData.append('arg2', this.tracedata.Fisherman_input.Sf_origin)
          formData.append('arg3', this.tracedata.Fisherman_input.Sf_salvageTime)
          formData.append('arg4', this.tracedata.Fisherman_input.Sf_outOfWaterTime)
          formData.append('arg5', this.tracedata.Fisherman_input.Sf_fishermanName)
          //formData.append('image', this.tracedata.Fisherman_input.Sf_duration)
          break
        case '加工厂':
          formData.append('arg1', this.tracedata.Factory_input.Fac_productName)
          formData.append('arg2', this.tracedata.Factory_input.Fac_productionbatch)
          formData.append('arg3', this.tracedata.Factory_input.Fac_productionTime)
          formData.append('arg4', this.tracedata.Factory_input.Fac_factoryName)
          formData.append('arg5', this.tracedata.Factory_input.Fac_contactNumber)
          break
        case '物流司机信息':
          formData.append('arg1', this.tracedata.Driver_input.Dr_name)
          formData.append('arg2', this.tracedata.Driver_input.Dr_age)
          formData.append('arg3', this.tracedata.Driver_input.Dr_phone)
          formData.append('arg4', this.tracedata.Driver_input.Dr_carNumber)
          formData.append('arg5', this.tracedata.Driver_input.Dr_transport)
          break
        case 'keyword_29':
          formData.append('arg1', this.tracedata.Shop_input.Sh_storeTime)
          formData.append('arg2', this.tracedata.Shop_input.Sh_sellTime)
          formData.append('arg3', this.tracedata.Shop_input.Sh_shopName)
          formData.append('arg4', this.tracedata.Shop_input.Sh_shopAddress)
          formData.append('arg5', this.tracedata.Shop_input.Sh_shopPhone)
          break
      }
      if (this.fileList.length > 0) {
        const reader = new FileReader();
        reader.readAsDataURL(this.fileList[0].raw);
        reader.onload = () => {
          formData.append('image', reader.result);
          console.log(formData)
          uplink(formData).then(res => {
            if (res.code === 200) {
              loading.close()
              this.$message({
                message: '上链成功，交易ID：' + res.txid + '\n溯源码：' + res.traceability_code,
                type: 'success'
              })
              this.resetForm()
            } else {
              loading.close()
              this.$message({
                message: '上链失败',
                type: 'error'
              })
            }
          }).catch(err => {
            loading.close()
            console.log(err)
          })
        };
      } else {
        uplink(formData).then(res => {
          if (res.code === 200) {
            loading.close()
            this.$message({
              message: '上链成功，交易ID：' + res.txid + '\n溯源码：' + res.traceability_code,
              type: 'success'
            })
            this.resetForm()
          } else {
            loading.close()
            this.$message({
              message: '上链失败',
              type: 'error'
            })
          }
        }).catch(err => {
          loading.close()
          console.log(err)
        })
      }
    },
    resetForm() {
      this.tracedata = {
        traceability_code: '',
        Fisherman_input: {
          Sf_seafoodName: '',
          Sf_origin: '',
          Sf_salvageTime: '',
          Sf_outOfWaterTime: '',
          Sf_fishermanName: ''
        },
        Factory_input: {
          Fac_productName: '',
          Fac_productionbatch: '',
          Fac_productionTime: '',
          Fac_factoryName: '',
          Fac_contactNumber: ''
        },
        Driver_input: {
          Dr_name: '',
          Dr_age: '',
          Dr_phone: '',
          Dr_carNumber: '',
          Dr_transport: ''
        },
        Shop_input: {
          Sh_storeTime: '',
          Sh_sellTime: '',
          Sh_shopName: '',
          Sh_shopAddress: '',
          Sh_shopPhone: ''
        }
      }
      this.fileList = []
    }
  }
}
</script>

<style lang="scss" scoped>
.uplink-container {
  padding: 30px;
  min-height: 100vh;

  &.fisherman-bg {
    background: linear-gradient(135deg, #1e3c72, #2a5298);
  }

  &.factory-bg {
    background: linear-gradient(135deg, #434343, #000000);
  }

  &.driver-bg {
    background: linear-gradient(135deg, #2c3e50, #3498db);
  }

  &.shop-bg {
    background: linear-gradient(135deg, #134e5e, #71b280);
  }

  &.default-bg {
    background: linear-gradient(135deg, #1a2a6c, #b21f1f, #fdbb2d);
  }

  .user-info-card {
    margin-bottom: 25px;
    border-radius: 15px;
    background: rgba(255, 255, 255, 0.95);
    backdrop-filter: blur(10px);

    .user-info {
      padding: 20px;

      .info-header {
        display: flex;
        justify-content: space-between;
        align-items: center;
        margin-bottom: 20px;
        border-bottom: 2px solid #eee;
        padding-bottom: 15px;

        h2 {
          color: #2c3e50;
          margin: 0;
          font-size: 24px;
          font-weight: 600;
        }

        .role-icon {
          i {
            font-size: 32px;
            color: #16a085;
          }
        }
      }

      .info-content {
        display: flex;
        align-items: center;

        .info-item {
          display: flex;
          align-items: center;

          i {
            margin-right: 10px;
            font-size: 18px;
            color: #16a085;
          }

          .user-text, .role-text {
            font-size: 16px;
            color: #2c3e50;
            letter-spacing: 0.5px;

            strong {
              font-weight: 600;
              color: #16a085;
              margin-left: 5px;
              font-size: 17px;
            }
          }
        }

        .el-divider {
          margin: 0 30px;
          height: 25px;
        }
      }
    }
  }

  .main-form-card {
    border-radius: 15px;
    background: rgba(255, 255, 255, 0.95);
    backdrop-filter: blur(10px);
    box-shadow: 0 8px 20px rgba(0,0,0,0.1);

    .role-form-section {
      background: rgba(236, 240, 241, 0.5);
      border-radius: 12px;
      padding: 25px;
      margin-bottom: 25px;
      border: 1px solid rgba(189, 195, 199, 0.3);
      transition: all 0.3s ease;

      &:hover {
        transform: translateY(-2px);
        box-shadow: 0 5px 15px rgba(0,0,0,0.1);
      }

      h3 {
        color: #2c3e50;
        margin: 0 0 25px 0;
        font-size: 20px;
        font-weight: 600;
        display: flex;
        align-items: center;
        justify-content: space-between;

        i {
          color: #16a085;
          font-size: 24px;

          &.role-specific-icon {
            font-size: 28px;
            color: #2980b9;
          }
        }
      }
    }

    .upload-section {
      background: rgba(236, 240, 241, 0.5);
      border-radius: 12px;
      padding: 25px;
      margin-bottom: 25px;
      border: 1px solid rgba(189, 195, 199, 0.3);
      text-align: center;
      transition: all 0.3s ease;

      &:hover {
        transform: translateY(-2px);
        box-shadow: 0 5px 15px rgba(0,0,0,0.1);
      }

      h3 {
        color: #2c3e50;
        margin: 0 0 25px 0;
        font-size: 20px;
        font-weight: 600;
        display: flex;
        align-items: center;
        justify-content: center;

        i {
          margin-right: 10px;
          color: #16a085;
          font-size: 24px;
        }
      }

      .upload-area {
        display: flex;
        justify-content: center;
        margin-bottom: 20px;

        ::v-deep .el-upload--picture-card {
          width: 180px;
          height: 180px;
          line-height: 180px;
          border-radius: 12px;
          border: none;
          background: transparent;
          transition: all 0.3s ease;

          &:hover {
            border-color: #409EFF;
            color: #409EFF;
            transform: scale(1.02);
          }
        }

        .upload-icon {
          font-size: 32px;
          color: #8c939d;
          margin-bottom: 10px;
        }

        .upload-text {
          display: flex;
          flex-direction: column;
          align-items: center;
          color: #606266;
          font-size: 14px;

          .upload-hint {
            margin-top: 8px;
            font-size: 12px;
            color: #909399;
          }
        }
      }

      .upload-tip {
        display: flex;
        align-items: center;
        justify-content: center;
        color: #909399;
        font-size: 13px;

        i {
          margin-right: 8px;
          font-size: 16px;
          color: #E6A23C;
        }
      }

      ::v-deep .el-upload-list--picture-card {
        display: flex;
        flex-wrap: wrap;
        justify-content: center;
        gap: 15px;

        .el-upload-list__item {
          width: 160px;
          height: 160px;
          border-radius: 8px;
          transition: all 0.3s ease;

          &:hover {
            transform: translateY(-3px);
          }
        }
      }
    }

    .trace-code-item {
      max-width: 500px;
    }

    .el-input {
      .el-input__inner {
        border-radius: 8px;
        border: 1px solid #dcdfe6;
        transition: all 0.3s ease;

        &:focus {
          border-color: #16a085;
          box-shadow: 0 0 0 2px rgba(22, 160, 133, 0.2);
        }
      }
    }

    .form-footer {
      margin-top: 40px;
      text-align: center;

      .el-button {
        padding: 12px 35px;
        font-size: 15px;
        border-radius: 8px;
        background: #16a085;
        border-color: #16a085;
        transition: all 0.3s ease;

        &:hover {
          background: #1abc9c;
          border-color: #1abc9c;
          transform: translateY(-2px);
          box-shadow: 0 5px 15px rgba(22, 160, 133, 0.3);
        }
      }

      .el-alert {
        margin-bottom: 25px;
        border-radius: 8px;
      }
    }
  }
}
</style>
