<template>
  <div v-show="loading">
    <div class="fullview">
      <div class="loading-spacer"></div>
      <vue-loading 
        type="bubbles"
        color="#BEDFC2"
        :size="{ width: '300px', height: '300px' }"
        >
      </vue-loading>
    </div>
  </div>
</template>

<script>
import { VueLoading } from 'vue-loading-template'

export default {
  name: 'loading',
  components: {
    VueLoading,
  },
  data(){
    return{
      loading: true
    }
  }

}
</script>

<style>
.fullview {
  width: 100%;
  height: 100%;
  background: #fefefe;
  position: fixed;
  top: 0;
  left: 0;
}
.loading-spacer {
  height: 20%;
}
</style>