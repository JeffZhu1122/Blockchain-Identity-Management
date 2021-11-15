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

    <el-table class="data_table" ref="filterTable" :data="tables.slice((currentPage-1)*pagesize,currentPage*pagesize)" >
      <el-table-column prop="recordId" label="ID" sortable >
      </el-table-column>
      <el-table-column prop="objectOfSale" label="People ID" sortable>
      </el-table-column>
      <el-table-column prop="agentId" label="Agent ID" sortable>
      </el-table-column>
      <el-table-column prop="branchId" label="Branch ID" sortable>
      </el-table-column>
      <el-table-column prop="from" label="From" sortable>
      </el-table-column>
      <el-table-column prop="fronDate" label="From Date" sortable>
      </el-table-column>
      <el-table-column prop="to" label="To" sortable>
      </el-table-column>
      <el-table-column prop="toDate" label="To Date" sortable>
      </el-table-column>
      <el-table-column prop="createTime" label="Checkin Date" sortable>
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


  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { querySellingList, createSellingByBuy, updateSelling } from '@/api/selling'

export default {
  name: 'AllSelling',
  data() {
    return {
      search: '',
      total: 0,
      pagesize:10,
      currentPage:1,
      loading: true,
      sellingList: []
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
          return  this.sellingList.filter(function(dataNews){
            return Object.keys(dataNews).some(function(key){
              return String(dataNews[key]).toLowerCase().indexOf(search) > -1
            })
          })
        }
        return this.sellingList
      }
  },
  created() {
    querySellingList().then(response => {
      if (response !== null) {
        this.sellingList = response
      }
      this.loading = false
    }).catch(_ => {
      this.loading = false
    })
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
    createSellingByBuy(item) {
      this.$confirm('是否立即购买?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'success'
      }).then(() => {
        this.loading = true
        createSellingByBuy({
          buyer: this.accountId,
          objectOfSale: item.objectOfSale,
          seller: item.seller
        }).then(response => {
          this.loading = false
          if (response !== null) {
            this.$message({
              type: 'success',
              message: '购买成功!'
            })
          } else {
            this.$message({
              type: 'error',
              message: '购买失败!'
            })
          }
          setTimeout(() => {
            window.location.reload()
          }, 1000)
        }).catch(_ => {
          this.loading = false
        })
      }).catch(() => {
        this.loading = false
        this.$message({
          type: 'info',
          message: '已取消购买'
        })
      })
    },
    updateSelling(item, type) {
      let tip = ''
      if (type === 'done') {
        tip = '确认收款'
      } else {
        tip = '取消操作'
      }
      this.$confirm('是否要' + tip + '?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'success'
      }).then(() => {
        this.loading = true
        updateSelling({
          buyer: item.buyer,
          objectOfSale: item.objectOfSale,
          seller: item.seller,
          status: type
        }).then(response => {
          this.loading = false
          if (response !== null) {
            this.$message({
              type: 'success',
              message: tip + '操作成功!'
            })
          } else {
            this.$message({
              type: 'error',
              message: tip + '操作失败!'
            })
          }
          setTimeout(() => {
            window.location.reload()
          }, 1000)
        }).catch(_ => {
          this.loading = false
        })
      }).catch(() => {
        this.loading = false
        this.$message({
          type: 'info',
          message: '已取消' + tip
        })
      })
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

  .all-card {
    width: 280px;
    height: 380px;
    margin: 18px;
  }
</style>
