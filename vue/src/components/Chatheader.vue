<template>
    <div>
        <div class="chatheader">
            <router-link class="header__item" to="/"> ポートフォーリオ画面へ</router-link>
            <router-link class="header__item" :to="'/users'">ユーザー一覧へ</router-link>
            <p class="header__item" @click="logout">ログアウト</p>
        </div>
    </div>
</template>

<script>
import router from '../router'
export default {
    methods: {
        logout() {
        localStorage.clear();
        router.replace('/login');
        }
    }
}
</script>

<style>

.chatheader{
  display: flex;
  justify-content: center;
  align-items: center;
  text-align: center;
  width: 100%;
  margin: 0 auto 20px;
}

.header__item{
  margin-left: 20px;
  cursor: pointer;
}


</style>