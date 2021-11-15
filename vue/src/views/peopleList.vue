<template>
  <div class="container">
    <el-row>
      <el-col :span=4>
        <div class="search-Box">
          <el-input placeholder="Keywords ..." icon="search"  class="search"  v-model="search"></el-input>
        </div>
      </el-col>
      <el-col :span=20><el-button style="float:right" @click="clearFilter">Clear Filter</el-button></el-col>
    </el-row>

    <el-table class="data_table" ref="filterTable" :data="tables.slice((currentPage-1)*pagesize,currentPage*pagesize)">
      <el-table-column prop="realEstateId" label="ID" sortable>
      </el-table-column>
      <el-table-column prop="firstname" label="Firstname" sortable >
      </el-table-column>
      <el-table-column prop="lastname" label="Lastname" sortable>
      </el-table-column>
      <el-table-column prop="dob" label="Date of Birth" sortable>
      </el-table-column>
      <el-table-column prop="status" label="Status" sortable>
      </el-table-column>
      <el-table-column prop="gender" label="Gender" sortable :filters="[{text: 'Male', value: 'male'}, {text: 'Female', value: 'famale'}]" :filter-method="filterHandler" column-key="gender" :formatter="formatLevel">
      </el-table-column>
      <el-table-column prop="proprietor" label="Agent" sortable >
      </el-table-column>
      <el-table-column prop="photo" label="Photo" >
        <template slot-scope="scope">
          <el-button v-if="scope.row.photo===''" type="primary" @click="showSingle('./assets/agent.jpg')">View</el-button>
          <el-button v-else type="primary" @click="showSingle(scope.row.photo)">View</el-button>
        </template>
      </el-table-column>
      <el-table-column prop="" label="Actions">
        <template slot-scope="scope">
          <el-button type="primary" @click="onPeopleDetail()">View Detail</el-button>
        </template>
      </el-table-column>

    </el-table>

    <div style="text-align: center;margin-top: 30px;">
      <el-pagination
        background
        layout="prev, pager, next, total"
        :total="total"
        :page-size="pagesize"
        @current-change="current_change">
      </el-pagination>
    </div>

    <vue-easy-lightbox
    :visible="visible"
    :imgs="imgs"
    :index="index"
    @hide="handleHide"
  ></vue-easy-lightbox>

  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { queryAccountList } from '@/api/account'
import { queryRealEstateList } from '@/api/realEstate'
import { createSelling } from '@/api/selling'
import { createDonating } from '@/api/donating'

export default {
  name: 'RealeState',
  data() {
    var checkArea = (rule, value, callback) => {
      if (value <= 0) {
        callback(new Error('必须大于0'))
      } else {
        callback()
      }
    }
    return {
      search: '',
      total: 0,
      pagesize:10,
      currentPage:1,

      imgs: '', // Img Url , string or Array of string
      visible: false,
      index: 0, // default: 0
      loading: true,
      loadingDialog: false,
      realEstateList: [],
      dialogCreateSelling: false,
      dialogCreateDonating: false,
      realForm: {
        price: 0,
        salePeriod: 0
      },
      rules: {
        price: [
          { validator: checkArea, trigger: 'blur' }
        ],
        salePeriod: [
          { validator: checkArea, trigger: 'blur' }
        ]
      },
      DonatingForm: {
        proprietor: ''
      },
      rulesDonating: {
        proprietor: [
          { required: true, message: '请选择业主', trigger: 'change' }
        ]
      },
      accountList: [],
      valItem: {}
    }
  },
  computed: {
    ...mapGetters([
      'accountId',
      'roles',
      'userName',
      'balance'
    ]),
    tables:function(){
        var search=this.search;
        if(search){
          return  this.realEstateList.filter(function(dataNews){
            return Object.keys(dataNews).some(function(key){
              return String(dataNews[key]).toLowerCase().indexOf(search) > -1
            })
          })
        }
        return this.realEstateList
      }
  },
  created() {
    if (this.roles[0] === 'admin'|| this.roles[0] === 'manager') {
      queryRealEstateList().then(response => {
        if (response !== null) {
          this.realEstateList = response

        }
        this.loading = false
      }).catch(_ => {
        this.loading = false
      })
    } else {
      queryRealEstateList({ proprietor: this.accountId }).then(response => {
        if (response !== null) {
          this.realEstateList = response
        }
        this.loading = false
      }).catch(_ => {
        this.loading = false
      })
    }
  },
  methods: {
    clearFilter() {
        this.$refs.filterTable.clearFilter();
        this.search=""
      },
    current_change(currentPage){
        this.currentPage = currentPage;
      },
    formatLevel(row, column) {
      return row.gender === 'famale' ? 'Female' : 'Male'
    },
    filterHandler(value, row) {
      return row.gender === value
    },
    showSingle(url) {
        this.imgs = url
        this.visible = true
      },
    handleHide() {
        this.visible = false
      },
    openDialog(item) {
      this.dialogCreateSelling = true
      this.valItem = item
    },
    onPeopleDetail(){
      this.$router.push({ path: '/people-detail' })
    },
    openDonatingDialog(item) {
      this.dialogCreateDonating = true
      this.valItem = item
      queryAccountList().then(response => {
        if (response !== null) {
          // 过滤掉管理员和当前用户
          this.accountList = response.filter(item =>
            item.userName !== '管理员' && item.accountId !== this.accountId
          )
        }
      })
    },
    createSelling(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$confirm('是否立即出售?', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'success'
          }).then(() => {
            this.loadingDialog = true
            createSelling({
              objectOfSale: this.valItem.realEstateId,
              seller: this.valItem.proprietor,
              price: this.realForm.price,
              salePeriod: this.realForm.salePeriod
            }).then(response => {
              this.loadingDialog = false
              this.dialogCreateSelling = false
              if (response !== null) {
                this.$message({
                  type: 'success',
                  message: '出售成功!'
                })
              } else {
                this.$message({
                  type: 'error',
                  message: '出售失败!'
                })
              }
              setTimeout(() => {
                window.location.reload()
              }, 1000)
            }).catch(_ => {
              this.loadingDialog = false
              this.dialogCreateSelling = false
            })
          }).catch(() => {
            this.loadingDialog = false
            this.dialogCreateSelling = false
            this.$message({
              type: 'info',
              message: '已取消出售'
            })
          })
        } else {
          return false
        }
      })
    },
    createDonating(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$confirm('是否立即捐赠?', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'success'
          }).then(() => {
            this.loadingDialog = true
            createDonating({
              objectOfDonating: this.valItem.realEstateId,
              donor: this.valItem.proprietor,
              grantee: this.DonatingForm.proprietor
            }).then(response => {
              this.loadingDialog = false
              this.dialogCreateDonating = false
              if (response !== null) {
                this.$message({
                  type: 'success',
                  message: '捐赠成功!'
                })
              } else {
                this.$message({
                  type: 'error',
                  message: '捐赠失败!'
                })
              }
              setTimeout(() => {
                window.location.reload()
              }, 1000)
            }).catch(_ => {
              this.loadingDialog = false
              this.dialogCreateDonating = false
            })
          }).catch(() => {
            this.loadingDialog = false
            this.dialogCreateDonating = false
            this.$message({
              type: 'info',
              message: '已取消捐赠'
            })
          })
        } else {
          return false
        }
      })
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
    },
    selectGet(accountId) {
      this.DonatingForm.proprietor = accountId
    }
  }
}

</script>

<style>
  .container{
    width: 100%;
    text-align: center;
    min-height: 100%;
    overflow: hidden;
  }
  .tag {
    float: left;
  }

  .item {
    font-size: 14px;
    margin-bottom: 18px;
    color: #999;
  }

  .clearfix:before,
  .clearfix:after {
    display: table;
  }
  .clearfix:after {
    clear: both
  }

  .realEstate-card {
    width: 280px;
    height: 340px;
    margin: 18px;
  }
</style>
