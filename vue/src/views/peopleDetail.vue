<template>
  <div class="container">

<div class="row mb-3">
    <div class="col-lg-4">
        <div class="card mb-3">
            <div class="card-body text-center shadow">
              <div class="col" style="float:left"><button class="btn btn-primary" @click="onBack()"  type="button" style="margin: 0;">Back</button></div>
              <img class="rounded-circle mb-3 mt-4" src="assets/img/dogs/image2.jpeg" width="160" height="160">
              <div class="col" ><button class="btn btn-primary" type="button" style="margin: 0;">Button</button></div>
            </div>

        </div>
        <div class="card shadow mb-4">
            <div class="card-header py-3">
                <h6 class="text-primary fw-bold m-0">Projects</h6>
            </div>
            <div class="card-body">
                <h4 class="small fw-bold">Server migration<span class="float-end">20%</span></h4>
                <div class="progress progress-sm mb-3">
                    <div class="progress-bar bg-danger" aria-valuenow="20" aria-valuemin="0" aria-valuemax="100" style="width: 20%;"><span class="visually-hidden">20%</span></div>
                </div>
                <h4 class="small fw-bold">Sales tracking<span class="float-end">40%</span></h4>
                <div class="progress progress-sm mb-3">
                    <div class="progress-bar bg-warning" aria-valuenow="40" aria-valuemin="0" aria-valuemax="100" style="width: 40%;"><span class="visually-hidden">40%</span></div>
                </div>
                <h4 class="small fw-bold">Customer Database<span class="float-end">60%</span></h4>
                <div class="progress progress-sm mb-3">
                    <div class="progress-bar bg-primary" aria-valuenow="60" aria-valuemin="0" aria-valuemax="100" style="width: 60%;"><span class="visually-hidden">60%</span></div>
                </div>
                <h4 class="small fw-bold">Payout Details<span class="float-end">80%</span></h4>
                <div class="progress progress-sm mb-3">
                    <div class="progress-bar bg-info" aria-valuenow="80" aria-valuemin="0" aria-valuemax="100" style="width: 80%;"><span class="visually-hidden">80%</span></div>
                </div>
                <h4 class="small fw-bold">Account setup<span class="float-end">Complete!</span></h4>
                <div class="progress progress-sm mb-3">
                    <div class="progress-bar bg-success" aria-valuenow="100" aria-valuemin="0" aria-valuemax="100" style="width: 100%;"><span class="visually-hidden">100%</span></div>
                </div>
            </div>
        </div>
    </div>
    <div class="col-lg-8">
        <div class="row mb-3 d-none">
            <div class="col">
                <div class="card textwhite bg-primary text-white shadow">
                    <div class="card-body">
                        <div class="row mb-2">
                            <div class="col">
                                <p class="m-0">Peformance</p>
                                <p class="m-0"><strong>65.2%</strong></p>
                            </div>
                            <div class="col-auto"><i class="fas fa-rocket fa-2x"></i></div>
                        </div>
                        <p class="text-white-50 small m-0"><i class="fas fa-arrow-up"></i>&nbsp;5% since last month</p>
                    </div>
                </div>
            </div>
            <div class="col">
                <div class="card textwhite bg-success text-white shadow">
                    <div class="card-body">
                        <div class="row mb-2">
                            <div class="col">
                                <p class="m-0">Peformance</p>
                                <p class="m-0"><strong>65.2%</strong></p>
                            </div>
                            <div class="col-auto"><i class="fas fa-rocket fa-2x"></i></div>
                        </div>
                        <p class="text-white-50 small m-0"><i class="fas fa-arrow-up"></i>&nbsp;5% since last month</p>
                    </div>
                </div>
            </div>
        </div>
        <div class="row">
            <div class="col">
                <div class="card shadow mb-3">
                    <div class="card-header py-3">
                        <p class="text-primary m-0 fw-bold">Information</p>
                    </div>
                    <div class="card-body" style="height: 517px;">
                        <form>
                            <div class="row">
                                <div class="col">
                                    <div class="mb-3"><label class="form-label" for="username"><strong>Username</strong></label>
                                        <h1>Heading</h1>
                                    </div>
                                </div>
                                <div class="col">
                                    <div class="mb-3"><label class="form-label" for="email"><strong>Email Address</strong></label>
                                        <h1>Heading</h1>
                                    </div>
                                </div>
                            </div>
                            <div class="row">
                                <div class="col">
                                    <div class="mb-3"><label class="form-label" for="first_name"><strong>First Name</strong></label>
                                        <h1>Heading</h1>
                                    </div>
                                </div>
                                <div class="col">
                                    <div class="mb-3"><label class="form-label" for="last_name"><strong>Last Name</strong></label>
                                        <h1>Heading</h1>
                                    </div>
                                </div>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>
<div class="card shadow mb-5">
    <div class="card-header py-3">
        <p class="text-primary m-0 fw-bold">Related Records</p>
    </div>
    <div class="card-body">
        <div class="row">
            <div class="col">
                <ul>
                    <li>Item 1</li>
                    <li>Item 2</li>
                    <li>Item 3</li>
                    <li>Item 4</li>
                </ul>
            </div>
        </div>
    </div>
</div>
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
    ])
  },
  created() {
    if (this.roles[0] === 'admin') {
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
    onBack(){
      this.$router.go(-1)
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

<style scoped>
  @import '../assets/static/bootstrap/css/bootstrap.min.css';
  @import '../assets/static/fonts/fontawesome-all.min.css';
  @import 'https://fonts.googleapis.com/css?family=Nunito:200,200i,300,300i,400,400i,600,600i,700,700i,800,800i,900,900i';

  .container{
    width: 100%;
    text-align: center;
    min-height: 100%;
    overflow: hidden;
    margin-top: 50px;
  }

  @media (min-width: 1400px){
    .container {
        max-width: 100%;
    }
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
