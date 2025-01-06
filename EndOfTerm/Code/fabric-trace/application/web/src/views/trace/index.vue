<template>
  <div v-loading="loading" element-loading-text="加载中...">
    <div class="trace-container">
      <div class="search-wrapper">
        <div class="search-area">
          <div class="search-item">
            <el-input
              v-model="traceCode"
              placeholder="请输入溯源码查询"
              class="custom-input"
              prefix-icon="el-icon-search"
            />
            <el-button
              type="primary"
              class="custom-button trace-button"
              @click="SeafoodInfo"
            >
              <i class="el-icon-search"></i>
              溯源码查询
            </el-button>
          </div>

          <div class="search-divider">
            <span>或</span>
          </div>

          <div class="search-item">
            <el-select
              v-model="searchField"
              placeholder="请选择搜索字段"
              class="custom-select"
            >
              <el-option label="水产品名称" value="sf_seafoodName" />
              <el-option label="捕捞区域" value="sf_origin" />
              <el-option label="渔民名称" value="sf_fishermanName" />
              <el-option label="加工厂名称" value="fac_factoryName" />
              <el-option label="销售点名称" value="sh_shopName" />
            </el-select>
            <el-input
              v-model="searchKeyword"
              placeholder="请输入搜索关键词"
              class="custom-input"
              prefix-icon="el-icon-edit"
            />
            <el-button
              type="primary"
              class="custom-button filter-button"
              @click="handleSearch"
            >
              <i class="el-icon-filter"></i>
              筛选查询
            </el-button>
          </div>
        </div>
      </div>
      <div class="tool-buttons">
        <el-button
          class="custom-button export-button"
          @click="exportPDF"
        >
          <i class="el-icon-download"></i>
          导出PDF
        </el-button>
        <el-button
          class="custom-button print-button"
          @click="printTrace"
        >
          <i class="el-icon-printer"></i>
          打印
        </el-button>
      </div>
      <div class="trace-list">
        <el-card
          v-for="item in filteredData"
          :key="item.traceability_code"
          class="trace-list-item"
          shadow="hover"
        >
          <div class="item-header" @click="toggleExpand(item)">
            <div class="item-info">
              <div class="info-row">
                <span class="label">溯源码：</span>
                <span class="value">{{ item.traceability_code }}</span>
              </div>
              <div class="info-row">
                <span class="label">水产品：</span>
                <span class="value">{{ item.fisherman_input.sf_seafoodName }}</span>
              </div>
              <div class="info-row">
                <span class="label">出水时间：</span>
                <span class="value">{{ formatTime(item.fisherman_input.sf_outOfWaterTime) }}</span>
              </div>
              <div class="info-row">
                <span class="label">当前状态：</span>
                <el-tag
                  :type="getStatusType(item)"
                  size="small"
                >
                  {{ getStatusText(item) }}
                </el-tag>
              </div>
            </div>
            <div class="item-action">
              <i :class="['el-icon-arrow-down', {'is-expanded': item.isExpanded}]"></i>
            </div>
          </div>

          <div class="progress-section">
            <el-steps :active="getActiveStep(item)" finish-status="success" align-center>
              <el-step title="捕捞" description="已完成捕捞">
                <template slot="icon">
                  <i class="el-icon-goods"></i>
                </template>
              </el-step>
              <el-step title="加工" description="工厂加工">
                <template slot="icon">
                  <i class="el-icon-office-building"></i>
                </template>
              </el-step>
              <el-step title="物流" description="运输配送">
                <template slot="icon">
                  <i class="el-icon-truck"></i>
                </template>
              </el-step>
              <el-step title="销售" description="终端销售">
                <template slot="icon">
                  <i class="el-icon-shopping-cart-full"></i>
                </template>
              </el-step>
            </el-steps>
          </div>

          <el-collapse-transition>
            <div v-show="item.isExpanded" class="item-detail">
              <div class="trace-detail">
                <el-card class="trace-card trace-card-product" shadow="hover">
                  <div class="card-header">
                    <i class="el-icon-goods"></i>
                    <span>水产品信息</span>
                  </div>
                  <div class="card-content">
                    <div class="info-section">
                      <el-form label-position="left" inline class="trace-form">
                        <el-form-item label="水产品名称：">
                          <span>{{ item.fisherman_input.sf_seafoodName || '暂无数据' }}</span>
                        </el-form-item>
                        <el-form-item label="捕捞区域：">
                          <span>{{ item.fisherman_input.sf_origin || '暂无数据' }}</span>
                        </el-form-item>
                        <el-form-item label="捕捞起始时间：">
                          <span>{{ formatTime(item.fisherman_input.sf_salvageTime) }}</span>
                        </el-form-item>
                        <el-form-item label="出水时间：">
                          <span>{{ formatTime(item.fisherman_input.sf_outOfWaterTime) }}</span>
                        </el-form-item>
                        <el-form-item label="渔民名称：">
                          <span>{{ item.fisherman_input.sf_fishermanName || '暂无数据' }}</span>
                        </el-form-item>
                        <el-form-item label="区块链交易ID：">
                          <span>{{ item.fisherman_input.sf_txid || '暂无数据' }}</span>
                        </el-form-item>
                        <el-form-item label="区块链交易时间：">
                          <span>{{ formatTime(item.fisherman_input.sf_timestamp) }}</span>
                        </el-form-item>
                      </el-form>
                    </div>
                    <div class="image-section">
                      <div class="image-title">图片记录</div>
                      <el-image
                        v-if="item.fisherman_input.image"
                        :src="item.fisherman_input.image"
                        fit="cover"
                        :preview-src-list="[item.fisherman_input.image]">
                        <div slot="error" class="image-slot">
                          <i class="el-icon-picture-outline"></i>
                        </div>
                      </el-image>
                      <div v-else class="no-image">
                        <i class="el-icon-picture-outline"></i>
                        <span>暂无图片</span>
                      </div>
                    </div>
                  </div>
                </el-card>

                <el-card class="trace-card trace-card-factory" shadow="hover">
                  <div class="card-header">
                    <i class="el-icon-office-building"></i>
                    <span>加工厂信息</span>
                  </div>
                  <div class="card-content">
                    <div class="info-section">
                      <el-form label-position="left" inline class="trace-form">
                        <el-form-item label="水产品商品名称：">
                          <span>{{ item.factory_input.fac_productName || '暂无数据' }}</span>
                        </el-form-item>
                        <el-form-item label="加工批次：">
                          <span>{{ item.factory_input.fac_productionbatch || '暂无数据' }}</span>
                        </el-form-item>
                        <el-form-item label="加工时间：">
                          <span>{{ formatTime(item.factory_input.fac_productionTime) }}</span>
                        </el-form-item>
                        <el-form-item label="加工厂名称与地址：">
                          <span>{{ item.factory_input.fac_factoryName || '暂无数据' }}</span>
                        </el-form-item>
                        <el-form-item label="加工厂电话：">
                          <span>{{ item.factory_input.fac_contactNumber || '暂无数据' }}</span>
                        </el-form-item>
                        <el-form-item label="区块链交易ID：">
                          <span>{{ item.factory_input.fac_txid || '暂无数据' }}</span>
                        </el-form-item>
                        <el-form-item label="区块链交易时间：">
                          <span>{{ formatTime(item.factory_input.fac_timestamp) }}</span>
                        </el-form-item>
                      </el-form>
                    </div>
                    <div class="image-section">
                      <div class="image-title">图片记录</div>
                      <el-image
                        v-if="item.factory_input.image"
                        :src="item.factory_input.image"
                        fit="cover"
                        :preview-src-list="[item.factory_input.image]">
                        <div slot="error" class="image-slot">
                          <i class="el-icon-picture-outline"></i>
                        </div>
                      </el-image>
                      <div v-else class="no-image">
                        <i class="el-icon-picture-outline"></i>
                        <span>暂无图片</span>
                      </div>
                    </div>
                  </div>
                </el-card>

                <el-card class="trace-card trace-card-logistics" shadow="hover">
                  <div class="card-header">
                    <i class="el-icon-truck"></i>
                    <span>物流与冷链轨迹信息</span>
                  </div>
                  <div class="card-content">
                    <div class="info-section">
                      <el-form label-position="left" inline class="trace-form">
                        <el-form-item label="物流司机姓名：">
                          <span>{{ item.driver_input.dr_name || '暂无数据' }}</span>
                        </el-form-item>
                        <el-form-item label="冷链车厢温度：">
                          <span>{{ item.driver_input.dr_temperature || '暂无数据' }}</span>
                        </el-form-item>
                        <el-form-item label="物流司机联系电话：">
                          <span>{{ item.driver_input.dr_phone || '暂无数据' }}</span>
                        </el-form-item>
                        <el-form-item label="运输车辆车牌号：">
                          <span>{{ item.driver_input.dr_carNumber || '暂无数据' }}</span>
                        </el-form-item>
                        <el-form-item label="运输与冷链记录：">
                          <span>{{ item.driver_input.dr_transport || '暂无数据' }}</span>
                        </el-form-item>
                        <el-form-item label="区块链交易ID：">
                          <span>{{ item.driver_input.dr_txid || '暂无数据' }}</span>
                        </el-form-item>
                        <el-form-item label="区块链交易时间：">
                          <span>{{ formatTime(item.driver_input.dr_timestamp) }}</span>
                        </el-form-item>
                      </el-form>
                    </div>
                    <div class="image-section">
                      <div class="image-title">图片记录</div>
                      <el-image
                        v-if="item.driver_input.image"
                        :src="item.driver_input.image"
                        fit="cover"
                        :preview-src-list="[item.driver_input.image]">
                        <div slot="error" class="image-slot">
                          <i class="el-icon-picture-outline"></i>
                        </div>
                      </el-image>
                      <div v-else class="no-image">
                        <i class="el-icon-picture-outline"></i>
                        <span>暂无图片</span>
                      </div>
                    </div>
                  </div>
                </el-card>

                <el-card class="trace-card trace-card-sales" shadow="hover">
                  <div class="card-header">
                    <i class="el-icon-shopping-cart-full"></i>
                    <span>销售终端信息</span>
                  </div>
                  <div class="card-content">
                    <div class="info-section">
                      <el-form label-position="left" inline class="trace-form">
                        <el-form-item label="销售终端">
                          <span>{{ item.shop_input.sh_storeTime || '暂无数据' }}</span>
                        </el-form-item>
                        <el-form-item label="销售时间：">
                          <span>{{ item.shop_input.sh_sellTime || '暂无数据' }}</span>
                        </el-form-item>
                        <el-form-item label="销售点名称：">
                          <span>{{ item.shop_input.sh_shopName || '暂无数据' }}</span>
                        </el-form-item>
                        <el-form-item label="销售点位置：">
                          <span>{{ item.shop_input.sh_shopAddress || '暂无数据' }}</span>
                        </el-form-item>
                        <el-form-item label="销售点电话：">
                          <span>{{ item.shop_input.sh_shopPhone || '暂无数据' }}</span>
                        </el-form-item>
                        <el-form-item label="区块链交易ID：">
                          <span>{{ item.shop_input.sh_txid || '暂无数据' }}</span>
                        </el-form-item>
                        <el-form-item label="区块链交易时间：">
                          <span>{{ formatTime(item.shop_input.sh_timestamp) }}</span>
                        </el-form-item>
                      </el-form>
                    </div>
                    <div class="image-section">
                      <div class="image-title">图片记录</div>
                      <el-image
                        v-if="item.shop_input.image"
                        :src="item.shop_input.image"
                        fit="cover"
                        :preview-src-list="[item.shop_input.image]">
                        <div slot="error" class="image-slot">
                          <i class="el-icon-picture-outline"></i>
                        </div>
                      </el-image>
                      <div v-else class="no-image">
                        <i class="el-icon-picture-outline"></i>
                        <span>暂无图片</span>
                      </div>
                    </div>
                  </div>
                </el-card>
              </div>
            </div>
          </el-collapse-transition>
        </el-card>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { getSeafoodInfo, getSeafoodList, getAllSeafoodInfo, getSeafoodHistory } from '@/api/trace'
import html2Canvas from 'html2canvas'
import JsPDF from 'jspdf'

export default {
  name: 'Trace',
  data() {
    return {
      tracedata: [],
      loading: false,
      traceCode: '',
      searchField: '',
      searchKeyword: '',
    }
  },
  computed: {
    ...mapGetters([
      'name',
      'userType'
    ]),
    filteredData() {
      if (!this.searchField || !this.searchKeyword) {
        return this.tracedata
      }

      return this.tracedata.filter(item => {
        let value = ''
        if (this.searchField.startsWith('sf_')) {
          value = item.fisherman_input[this.searchField]
        } else if (this.searchField.startsWith('fac_')) {
          value = item.factory_input[this.searchField]
        } else if (this.searchField.startsWith('sh_')) {
          value = item.shop_input[this.searchField]
        }

        return value && value.toLowerCase().includes(this.searchKeyword.toLowerCase())
      })
    }
  },
  created() {
    getAllSeafoodInfo().then(res => {
      this.tracedata = JSON.parse(res.data).map(item => ({
        ...item,
        isExpanded: false
      }))
    })
  },
  methods: {
    SeafoodHistory() {
      getSeafoodHistory().then(res => {
        // console.log(res)
      })
    },
    SeafoodInfo() {
      var formData = new FormData()
      formData.append('traceability_code', this.traceCode)
      getSeafoodInfo(formData).then(res => {
        if (res.code === 200) {
          this.tracedata = []
          this.tracedata[0] = {
            ...JSON.parse(res.data),
            isExpanded: false
          }
          return
        } else {
          this.$message.error(res.message)
        }
      })
    },
    handleSearch() {
      if (!this.searchField) {
        this.$message.warning('请选择搜索字段')
        return
      }
      if (!this.searchKeyword) {
        this.$message.warning('请输入搜索关键词')
        return
      }
    },
    getStatusType(item) {
      if (item.shop_input.sh_sellTime) {
        return 'success'
      } else if (item.driver_input.dr_transport) {
        return 'warning'
      } else if (item.factory_input.fac_productionTime) {
        return 'primary'
      }
      return 'info'
    },
    getStatusText(item) {
      if (item.shop_input.sh_sellTime) {
        return '已售出'
      } else if (item.driver_input.dr_transport) {
        return '运输中'
      } else if (item.factory_input.fac_productionTime) {
        return '加工完成'
      }
      return '已捕捞'
    },
    toggleExpand(item) {
      this.$set(item, 'isExpanded', !item.isExpanded)
    },
    getTimelineType(stage, item) {
      if (stage === 0) return 'primary'
      if (stage === 1 && item.factory_input.fac_productionTime) return 'primary'
      if (stage === 2 && item.driver_input.dr_transport) return 'primary'
      if (stage === 3 && item.shop_input.sh_sellTime) return 'primary'
      return 'info'
    },
    getTimelineIcon(stage) {
      const icons = [
        'el-icon-goods',
        'el-icon-office-building',
        'el-icon-truck',
        'el-icon-shopping-cart-full'
      ]
      return icons[stage]
    },
    async exportPDF() {
      if (!this.tracedata.length) {
        this.$message.warning('暂无数据可导出')
        return
      }

      try {
        this.$message({
          message: '正在生成PDF，请稍候...',
          type: 'info'
        })

        // 如果有展开的记录，就导出展开的记录，否则导出第一条记录
        const expandedItem = this.tracedata.find(item => item.isExpanded)
        if (expandedItem) {
          // 确保记录展开
          await this.$nextTick()
          const element = document.querySelector('.item-detail')
          if (!element) {
            throw new Error('未找到要导出的内容')
          }

          const canvas = await html2Canvas(element, {
            allowTaint: true,
            useCORS: true,
            scale: 2 // 提高清晰度
          })

          const contentWidth = canvas.width
          const contentHeight = canvas.height
          const pageHeight = (contentWidth / 592.28) * 841.89
          const leftHeight = contentHeight
          const position = 0
          const imgWidth = 592.28
          const imgHeight = (592.28 / contentWidth) * contentHeight

          const pdf = new JsPDF('p', 'pt', 'a4')
          if (leftHeight < pageHeight) {
            pdf.addImage(canvas.toDataURL('image/jpeg', 1.0), 'JPEG', 0, 0, imgWidth, imgHeight)
          } else {
            while (leftHeight > 0) {
              pdf.addImage(canvas.toDataURL('image/jpeg', 1.0), 'JPEG', 0, position, imgWidth, imgHeight)
              leftHeight -= pageHeight
              position -= 841.89
              if (leftHeight > 0) {
                pdf.addPage()
              }
            }
          }

          pdf.save(`溯源信息_${expandedItem.traceability_code}.pdf`)
          this.$message.success('PDF导出成功')
        } else {
          this.$message.warning('请先展开要导出的记录')
        }
      } catch (error) {
        console.error('PDF导出失败:', error)
        this.$message.error('PDF导出失败，请重试')
      }
    },

    async printTrace() {
      if (!this.tracedata.length) {
        this.$message.warning('暂无数据可打印')
        return
      }

      const expandedItem = this.tracedata.find(item => item.isExpanded)
      if (!expandedItem) {
        this.$message.warning('请先展开要打印的记录')
        return
      }

      try {
        await this.$nextTick()
        const element = document.querySelector('.item-detail')
        if (!element) {
          throw new Error('未找到要打印的内容')
        }

        const printWindow = window.open('', '_blank')
        printWindow.document.write(`
          <!DOCTYPE html>
          <html>
            <head>
              <title>溯源信息_${expandedItem.traceability_code}</title>
              <style>
                body { font-family: Arial, sans-serif; }
                .trace-card {
                  page-break-inside: avoid;
                  margin-bottom: 20px;
                  border: 1px solid #EBEEF5;
                  padding: 20px;
                  border-radius: 8px;
                }
                .card-header {
                  font-size: 18px;
                  font-weight: bold;
                  margin-bottom: 15px;
                }
                .info-section {
                  display: grid;
                  grid-template-columns: repeat(2, 1fr);
                  gap: 10px;
                }
                .el-form-item {
                  margin-bottom: 10px;
                }
                @media print {
                  .trace-card {
                    break-inside: avoid;
                  }
                }
              </style>
            </head>
            <body>
              ${element.innerHTML}
            </body>
          </html>
        `)

        printWindow.document.close()
        printWindow.focus()

        // 等待图片加载完成后打印
        setTimeout(() => {
          printWindow.print()
          printWindow.close()
        }, 500)
      } catch (error) {
        console.error('打印失败:', error)
        this.$message.error('打印失败，请重试')
      }
    },
    getActiveStep(item) {
      if (item.shop_input.sh_sellTime) {
        return 4
      } else if (item.driver_input.dr_transport) {
        return 3
      } else if (item.factory_input.fac_productionTime) {
        return 2
      }
      return 1
    },
    getStepDescription(item, step) {
      const descriptions = {
        1: `捕捞时间: ${item.fisherman_input.sf_outOfWaterTime || '未知'}`,
        2: `加工时间: ${item.factory_input.fac_productionTime || '未开始'}`,
        3: `运输时间: ${item.driver_input.dr_transport || '未开始'}`,
        4: `销售时间: ${item.shop_input.sh_sellTime || '未开始'}`
      }
      return descriptions[step]
    },
    formatTime(time) {
      if (!time) return '暂无数据'
      try {
        const date = new Date(time)
        const year = date.getFullYear()
        const month = String(date.getMonth() + 1).padStart(2, '0')
        const day = String(date.getDate()).padStart(2, '0')
        const hours = String(date.getHours()).padStart(2, '0')
        const minutes = String(date.getMinutes()).padStart(2, '0')
        const seconds = String(date.getSeconds()).padStart(2, '0')

        return `${year}.${month}.${day} ${hours}:${minutes}:${seconds}`
      } catch (error) {
        console.error('时间格式化失败:', error)
        return '暂无数据'
      }
    }
  }
}
</script>

<style lang="scss" scoped>

.demo-table-expand {
    font-size: 0;
  }
  .demo-table-expand label {
    width: 90px;
    color: #99a9bf;
  }
  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 0;
    width: 50%;
  }
.trace {
  &-container {
    margin: 30px;
  }
  &-text {
    font-size: 30px;
    line-height: 46px;
  }
}

.search-wrapper {
  display: flex;
  justify-content: center;
  padding: 40px 0;
  background: linear-gradient(to bottom, #f6f8fa, #ffffff);
}

.search-area {
  width: 800px;
  padding: 30px;
  background: white;
  border-radius: 12px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);


  .search-item {
    display: flex;
    align-items: center;
    gap: 15px;

    .custom-input {
      flex: 1;

      :deep(.el-input__inner) {
        border-radius: 8px;
        height: 42px;
        transition: all 0.3s;

        &:focus {
          box-shadow: 0 0 0 2px rgba(64, 158, 255, 0.2);
        }
      }
    }

    .custom-select {
      width: 180px;

      :deep(.el-input__inner) {
        border-radius: 8px;
        height: 42px;
      }
    }

    .custom-button {
      min-width: 120px;
      height: 42px;
      border-radius: 8px;
      font-weight: 500;
      letter-spacing: 1px;
      transition: all 0.3s;
      position: relative;
      overflow: hidden;

      i {
        margin-right: 6px;
      }

      &::before {
        content: '';
        position: absolute;
        top: 50%;
        left: 50%;
        width: 0;
        height: 0;
        background: rgba(255, 255, 255, 0.2);
        border-radius: 50%;
        transform: translate(-50%, -50%);
        transition: width 0.6s, height 0.6s;
      }

      &:hover {
        transform: translateY(-2px);
        box-shadow: 0 4px 12px rgba(64, 158, 255, 0.3);

        &::before {
          width: 300px;
          height: 300px;
        }
      }

      &:active {
        transform: translateY(1px);
      }
    }


    .trace-button {
      background: linear-gradient(45deg, #409EFF, #36D1DC);
      border: none;

      &:hover {
        background: linear-gradient(45deg, #36D1DC, #409EFF);
      }
    }

    .filter-button {
      background: linear-gradient(45deg, #67C23A, #36D1DC);
      border: none;

      &:hover {
        background: linear-gradient(45deg, #36D1DC, #67C23A);
      }
    }
  }

  .search-divider {
    display: flex;
    align-items: center;
    margin: 20px 0;
    color: #909399;

    &::before,
    &::after {
      content: '';
      flex: 1;
      height: 1px;
      background: #EBEEF5;
    }

    span {
      padding: 0 15px;
      font-size: 14px;
    }
  }
}

.trace-detail {
  padding: 20px;

  .trace-card {
    margin-bottom: 20px;
    border-radius: 8px;
    transition: all 0.3s;

    &:hover {
      transform: translateY(-5px);
      box-shadow: 0 6px 12px rgba(0,0,0,0.1);
    }

    .card-header {
      display: flex;
      align-items: center;
      margin-bottom: 20px;

      i {
        font-size: 24px;
        margin-right: 10px;
      }

      span {
        font-size: 18px;
        font-weight: 600;
      }
    }

    .card-content {
      display: flex;

      .info-section {
        flex: 2;
        padding-right: 20px;
        border-right: 1px solid #EBEEF5;
      }

      .image-section {
        flex: 1;
        padding-left: 20px;
        display: flex;
        flex-direction: column;
        align-items: center;

        .image-title {
          font-size: 16px;
          font-weight: 500;
          color: #606266;
          margin-bottom: 15px;
        }

        .el-image {
          width: 100%;
          height: 200px;
          border-radius: 4px;

          .image-slot {
            display: flex;
            justify-content: center;
            align-items: center;
            width: 100%;
            height: 100%;
            background: #f5f7fa;
            color: #909399;
            font-size: 30px;
          }
        }

        .no-image {
          width: 100%;
          height: 200px;
          background: #f5f7fa;
          border-radius: 4px;
          display: flex;
          flex-direction: column;
          justify-content: center;
          align-items: center;
          color: #909399;

          i {
            font-size: 40px;
            margin-bottom: 10px;
          }

          span {
            font-size: 14px;
          }
        }
      }
    }
  }

  .trace-card-product {
    .card-header {
      color: #67C23A;
    }
  }

  .trace-card-factory {
    .card-header {
      color: #409EFF;
    }
  }

  .trace-card-logistics {
    .card-header {
      color: #E6A23C;
    }
  }

  .trace-card-sales {
    .card-header {
      color: #909399;
    }
  }

  .trace-form {
    font-size: 14px;

    .el-form-item {
      margin-bottom: 10px;
      width: 33.33%;
      padding-right: 15px;

      label {
        color: #606266;
        font-weight: 500;
        min-width: 120px !important;
      }

      span {
        color: #303133;
        font-size: 13px;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
        max-width: 200px;
        display: inline-block;

        &:empty::before,
        &:contains('暂无数据') {
          content: '暂无数据';
          color: #909399;
          font-style: italic;
        }
      }
    }
  }
}

.trace-list {
  .trace-list-item {
    margin-bottom: 15px;
    border-radius: 8px;
    cursor: pointer;

    &:hover {
      .item-header {
        background-color: #f5f7fa;
      }
    }

    .item-header {
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding: 15px;
      transition: all 0.3s;
      border-radius: 8px;

      .item-info {
        flex: 1;
        display: grid;
        grid-template-columns: repeat(4, 1fr);
        gap: 20px;

        .info-row {
          display: flex;
          align-items: center;

          .label {
            color: #909399;
            margin-right: 8px;
          }

          .value {
            color: #303133;
            font-weight: 500;
          }
        }
      }

      .item-action {
        margin-left: 20px;

        i {
          font-size: 20px;
          color: #909399;
          transition: all 0.3s;

          &.is-expanded {
            transform: rotate(180deg);
          }
        }
      }
    }

    .item-detail {
      padding: 20px 15px;
      border-top: 1px solid #EBEEF5;
    }
  }
}

.el-collapse-transition {
  transition: 0.3s height ease-in-out, 0.3s padding-top ease-in-out, 0.3s padding-bottom ease-in-out;
}

.timeline-section {
  margin: 30px 0;
  padding: 20px;
  background: #f8f9fa;
  border-radius: 8px;

  .el-timeline {
    padding: 20px;
  }

  .el-timeline-item {
    h4 {
      color: #303133;
      margin: 0 0 10px;
    }

    p {
      color: #606266;
      margin: 5px 0;
      font-size: 14px;
    }
  }
}

.tool-buttons {
  display: flex;
  gap: 15px;
  margin: 20px 40px;
  justify-content: flex-end;

  .custom-button {
    min-width: 120px;
    height: 42px;
    border-radius: 8px;
    font-weight: 500;
    letter-spacing: 1px;
    transition: all 0.3s;
    position: relative;
    overflow: hidden;
    border: none;
    color: white;
    padding: 0 25px;

    i {
      margin-right: 8px;
      font-size: 16px;
    }

    &::before {
      content: '';
      position: absolute;
      top: 50%;
      left: 50%;
      width: 0;
      height: 0;
      background: rgba(255, 255, 255, 0.2);
      border-radius: 50%;
      transform: translate(-50%, -50%);
      transition: width 0.6s, height 0.6s;
    }

    &:hover {
      transform: translateY(-2px);
      box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);

      &::before {
        width: 300px;
        height: 300px;
      }
    }

    &:active {
      transform: translateY(1px);
    }
  }

  .export-button {
    background: linear-gradient(45deg, #FF9800, #F44336);

    &:hover {
      background: linear-gradient(45deg, #F44336, #FF9800);
    }
  }

  .print-button {
    background: linear-gradient(45deg, #9C27B0, #673AB7);

    &:hover {
      background: linear-gradient(45deg, #673AB7, #9C27B0);
    }
  }
}

.progress-section {
  padding: 30px 20px;
  margin: 20px 0;
  background: #f8f9fa;
  border-radius: 8px;
  transition: all 0.3s;

  &:hover {
    background: #fff;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
  }

  :deep(.el-steps) {
    .el-step__head {
      .el-step__icon {
        width: 40px;
        height: 40px;
        font-size: 20px;

        &.is-text {
          border-radius: 50%;
          border: 2px solid;
          transition: all 0.3s;

          .el-step__icon-inner {
            font-weight: bold;

            i {
              font-size: 18px;
            }
          }
        }
      }

      &.is-process {
        .el-step__icon {
          background: linear-gradient(45deg, #409EFF, #36D1DC);
          border-color: transparent;
          color: white;
          transform: scale(1.1);
          box-shadow: 0 4px 12px rgba(64, 158, 255, 0.3);
        }
      }

      &.is-success {
        .el-step__icon {
          background: linear-gradient(45deg, #67C23A, #36D1DC);
          border-color: transparent;
          color: white;
        }
      }

      &.is-wait {
        .el-step__icon {
          background: #f5f7fa;
          border-color: #dcdfe6;
          color: #c0c4cc;
        }
      }
    }

    .el-step__main {
      .el-step__title {
        font-size: 16px;
        font-weight: 500;

        &.is-process {
          color: #409EFF;
        }

        &.is-success {
          color: #67C23A;
        }
      }

      .el-step__description {
        font-size: 14px;

        &.is-process {
          color: #409EFF;
        }

        &.is-success {
          color: #67C23A;
        }
      }
    }
  }
}
</style>
