<template>
  <div>{{log}}</div>
</template>

<script>
import { mapActions } from 'vuex'
import { squarePreviewDiv } from '@/utils.js'
import axios from "axios";

export default {
  data() {
    return { log: '' }
  },
  created() {
    let options = {
      method: 'GET',
      url: `/admin/log`,
    }
    axios(options).then(res => {
      this.log = res
    })
  },
  mounted() {
  },
}
</script>
