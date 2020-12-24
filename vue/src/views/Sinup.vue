<template>
    <div class="sinup">
        <header class="sininheader">
            <router-link class="header__item" to="/"> ポートフォーリオ画面へ</router-link>
            <router-link class="header__item" to="/login">ログイン画面へ</router-link>
        </header>
        <h2 class="sinup__ttl">ユーザー登録</h2>
        <table class="sinup__form">
            <tr>
                <th>
                    <label for="name">ユーザー名</label>
                </th>
                <td>
                    <input class="sinup__input" type="name" id="name" v-model="username">
                </td>
            </tr>
            <tr>
                <th>
                    <label for="password">パスワード</label>
                </th>
                <td>
                    <input class="sinup__input" type="password" id="password" v-model="password" maxlength="8" pattern="^[0-9A-Za-z]+$">
                    <p class="sinup__password_txt">＊パスワードは４〜８桁で入力してください</p>
                </td>
            </tr>
        </table>

        <button @click="sinup">ログイン</button>


    </div>
</template>

<script>
import axios from 'axios';
import router from '../router';

export default {
    data() {
        return {
            username: "",
            password: ""
        }
    },
    methods: {
        sinup(){
            const params = new URLSearchParams();
            params.append('username', this.username);
            params.append('password', this.password);

            axios.post('/usersinup', params).then(response => {
                if(!response.data.errormessage) {
                    let user = response.data.user;
                    localStorage.setItem('username', user.Username);
                    localStorage.setItem('userID', user.ID);
                    router.push('/users');
                }else{
                    let errorMessage = response.data.errormessage;
                    errorMessage = errorMessage.join("");
                    alert(errorMessage);
                }
            })

            this.username = "";
            this.password = "";
        }
    }
}
</script>


<style scoped>

.sinup{
    height: 100vh;
    /* margin: 10vh auto; */
    text-align: center;
    /* background: #BEDFC2; */
}

.sinup__ttl{
    font-size: 32px;
}

.sinup__form{
    text-align: center;
    width: 50%;
    margin: 40px auto;
}

.sinup__form tr{
    margin-bottom: 10px;
}

.sinup__input{
  position: relative;
  display: block;
  /* width: 500px; */
  padding: 15px;
  border: 2px solid #aaa;
  border-radius: 5px;
  font-size: 16px;
  outline: none;
  text-align: left;
  margin: 10px;
}

.sinup__password_txt{
    font-size: 12px;
    padding-top: 10px;
    text-align: left;
}

.sininheader{
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