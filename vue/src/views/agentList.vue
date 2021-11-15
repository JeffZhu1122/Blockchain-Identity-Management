<template>
  <div class="app-container">
    <el-row>
      <el-col :span=4>
        <div class="search-Box">
          <el-input placeholder="Keywords ..." icon="search"  class="search"  v-model="search"></el-input>
        </div>
      </el-col>
      <el-col :span=20><el-button style="float:right" @click="clearFilter">Clear Filter</el-button></el-col>
    </el-row>


    <el-table ref="filterTable" class="data_table" :data="tables.slice((currentPage-1)*pagesize,currentPage*pagesize)">
      <el-table-column prop="accountId" label="ID" sortable>
      </el-table-column>
      <el-table-column prop="email" label="Email" sortable>
      </el-table-column>
      <el-table-column prop="phone" label="Phone" sortable>
      </el-table-column>
      <el-table-column prop="uname" label="Username" sortable>
      </el-table-column>
      <el-table-column prop="level" label="Level" sortable :filters="[{text: 'Admin', value: 0}, {text: 'Manager', value: 1}, {text: 'Agent', value: 2}]" :filter-method="filterHandler" column-key="level" :formatter="formatLevel">
      </el-table-column>
      <el-table-column prop="photo" label="Photo">
        <template slot-scope="scope">
          <el-button type="primary" @click="showSingle('https://dl.shimmertech.net/agent.jpg')">View</el-button>
        </template>
      </el-table-column>
      <el-table-column prop="" label="Actions">
        <template slot-scope="scope">
          <b-button @click="toggleModal(scope.row)">Manage</b-button>
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


    <!-- The modal -->
    <b-modal ref="my-modal" title="Edit Agent">
      <el-form ref="form" label-width="80px">
        <el-form-item label="Email">
          <el-input v-model="editForm.email"></el-input>
        </el-form-item>
        <el-form-item label="Phone">
          <el-input v-model="editForm.phone"></el-input>
        </el-form-item>
        <el-form-item label="Level">
          <el-select v-model="editForm.Level" placeholder="please select level">
            <el-option label="Admin" value=0></el-option>
            <el-option label="Manager" value=1></el-option>
            <el-option label="Agent" value=2></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="Photo">
          <el-upload
            class="upload-demo"
            drag
            action="https://jsonplaceholder.typicode.com/posts/"
            >
            <i class="el-icon-upload"></i>
            <div class="el-upload__text">Drop file here or <em>click to upload</em></div>
            <div class="el-upload__tip" slot="tip">jpg/png files with a size less than 2000kb</div>
          </el-upload>
        </el-form-item>

      </el-form>
    </b-modal>


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
import { createRealEstate } from '@/api/realEstate'

export default {
  name: 'AddRealeState',
  data() {
    var checkArea = (rule, value, callback) => {
      if (value <= 0) {
        callback(new Error('必须大于0'))
      } else {
        callback()
      }
    }
    return {
      total: 0,
      editForm:{
          email:'',
          phone:'',
          Level:'',
          photo:''
      },
      pagesize:10,
      currentPage:1,
      search: '',
      imgs: '', // Img Url , string or Array of string
      visible: false,
      index: 0, // default: 0
      ruleForm: {
        proprietor: '',
        totalArea: 0,
        livingSpace: 0
      },
      accountList: [],
      rules: {
        proprietor: [
          { required: true, message: '请选择业主', trigger: 'change' }
        ],
        totalArea: [
          { validator: checkArea, trigger: 'blur' }
        ],
        livingSpace: [
          { validator: checkArea, trigger: 'blur' }
        ]
      },
      loading: false
    }
  },
  computed: {
    ...mapGetters([
      'accountId'
    ]),
    tables:function(){
        var search=this.search;
        if(search){
          return  this.accountList.filter(function(dataNews){
            return Object.keys(dataNews).some(function(key){
              return String(dataNews[key]).toLowerCase().indexOf(search) > -1
            })
          })
        }
        return this.accountList
      }
  },
  created() {
    queryAccountList().then(response => {
      if (response !== null) {
        this.accountList = response
        this.total=response.length
        // .filter(item =>
        //   item.userName !== '管理员'
        // )
      }
    })
  },
  methods: {
    toggleModal(row) {
        this.editForm.email=row.email
        this.editForm.phone=row.phone
        this.editForm.level=row.level

        this.$refs['my-modal'].toggle('#toggle-btn')
      },
    clearFilter() {
        this.$refs.filterTable.clearFilter();
        this.search=""
      },
    current_change(currentPage){
        this.currentPage = currentPage;
      },
    formatLevel(row, column) {
      return row.level === 0 ? 'Admin' : row.level === 1 ? 'Manager' : 'Agent'
    },
    filterHandler(value, row) {
      return row.level === value
    },
    showSingle(url) {
        this.imgs = url
        this.visible = true
      },
    handleHide() {
        this.visible = false
      },
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          this.$confirm('是否立即创建?', '提示', {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'success'
          }).then(() => {
            this.loading = true
            createRealEstate({
              accountId: this.accountId,
              proprietor: this.ruleForm.proprietor,
              totalArea: this.ruleForm.totalArea,
              livingSpace: this.ruleForm.livingSpace
            }).then(response => {
              this.loading = false
              if (response !== null) {
                this.$message({
                  type: 'success',
                  message: '创建成功!'
                })
              } else {
                this.$message({
                  type: 'error',
                  message: '创建失败!'
                })
              }
            }).catch(_ => {
              this.loading = false
            })
          }).catch(() => {
            this.loading = false
            this.$message({
              type: 'info',
              message: '已取消创建'
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
      this.ruleForm.proprietor = accountId
    }
  }
}
</script>

<style scoped>
</style>
