
<template>
  <div class="app-container">
    <el-form ref="form" :model="form" label-width="120px">
      <el-form-item label="Firs name">
        <el-input v-model="form.firstname"></el-input>
      </el-form-item>
      <el-form-item label="Lastname">
        <el-input v-model="form.lastname"></el-input>
      </el-form-item>
      <el-form-item label="Gender">
        <el-select v-model="form.gender" placeholder="please select your gender">
          <el-option label="Male" value="male"></el-option>
          <el-option label="Female" value="female"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item label="Date of Birth">
          <el-date-picker type="date" placeholder="Pick a date" v-model="form.dob" style="width: 100%;"></el-date-picker>
      </el-form-item>
      <el-form-item label="State">
        <el-switch v-model="form.state"></el-switch>
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
      form: {
        firstname: '',
        lastname: '',
        gender: '',
        dob: '',
        state: '',
        photo: '',
        }



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
