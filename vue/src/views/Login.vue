<template>
    <div class="singin">
        <header class="sininheader">
            <router-link class="header__item" to="/"> ポートフォーリオ画面へ</router-link>
            <router-link class="header__item" to="/sinup">ユーザー登録画面へ</router-link>
        </header>
        <h2 class="singin__ttl">ログイン画面</h2>
        <table class="singin__form">
            <tr>
                <th>
                    <label for="name">ユーザー名</label>
                </th>
                <td>
                    <input class="singin__input" type="name" id="name" v-model="username">
                </td>
            </tr>
            <tr>
                <th>
                    <label for="password">パスワード</label>
                </th>
                <td>
                    <input class="singin__input" type="password" id="password" v-model="password" maxlength="8" pattern="^[0-9A-Za-z]+$">
                    <p class="singin__password_txt">＊パスワードは４〜８桁で入力してください</p>
                </td>
            </tr>
        </table>

        <button @click="singin">ログイン</button>


    </div>
</template>

<script>
import axios from 'axios';
import router from '../router.js'

export default {
    data() {
        return {
            username: "",
            password: ""
        }
    },
    methods: {
        singin(){
            const params = new URLSearchParams();
            params.append('username', this.username);
            params.append('password', this.password);

            axios.post('/usersinin', params).then(response => {
                if(!response.data.message) {
                    alert("ログインに失敗しました")
                }else{
                    let user = response.data.user;
                    localStorage.setItem('username', user.Username);
                    localStorage.setItem('userID', user.ID);
                    router.push('/users');
                }
            })

            this.username = "";
            this.password = "";
        }
    }
}
</script>


<style scoped>

.singin{
    height: 100vh;
    margin: 10vh auto;
    text-align: center;
    /* background: #BEDFC2; */
}

.singin__ttl{
    font-size: 32px;
}

.singin__form{
    text-align: center;
    width: 50%;
    margin: 40px auto;
}

.singin__form tr{
    margin-bottom: 10px;
}

.singin__input{
  position: relative;
  display: block;
  padding: 15px;
  border: 2px solid #aaa;
  border-radius: 5px;
  font-size: 16px;
  outline: none;
  text-align: left;
  margin: 10px;
}

.singin__password_txt{
    font-size: 12px;
    padding-top: 10px;
    text-align: left;
}

.sininheader{
    display: flex;
    justify-content: center;
    align-items: center;
    text-align: center;
    width: 80%;
    margin: 0 auto 20px;
}

.header__item{
    margin-left: 20px;
    cursor: pointer;
}

</style>