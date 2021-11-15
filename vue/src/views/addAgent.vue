
<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="Email">
        <el-input v-model="editForm.email"></el-input>
      </el-form-item>
      <el-form-item label="Password">
        <el-input v-model="editForm.password"></el-input>
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
          :on-preview="handlePreview"
          :on-remove="handleRemove"
          :file-list="fileList"
          >
          <i class="el-icon-upload"></i>
          <div class="el-upload__text">Drop file here or <em>click to upload</em></div>
          <div class="el-upload__tip" slot="tip">jpg/png files with a size less than 2000kb</div>
        </el-upload>
      </el-form-item>

    </el-form>



  </div>
</template>

<script>
import { mapGetters } from 'vuex'
import { queryAccountList } from '@/api/account'
import { createRealEstate } from '@/api/realEstate'

export default {
  name: 'AddRealeState',
  data() {

    return {
      fileList:[],
      editForm:{
          email:'',
          phone:'',
          Level:'',
          photo:'',
          password:''
      },



    }
  },
  computed: {

  },
  created() {

  },
  methods: {

  }
}
</script>

<style scoped>
</style>
