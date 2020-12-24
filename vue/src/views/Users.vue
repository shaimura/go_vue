<template>
    <div class="users">
        <Chatheader></Chatheader>
        <h2 class="users__ttl">ユーザー選択画面</h2>
        <div v-for="user in users" :key="user.ID">
            <table class="user__table">
                <tr>
                    <th>ユーザー名：</th>
                    <td>{{ user.Username }}</td>
                    <td>
                        <router-link @click="userchatroom(user.ID)" 
                                     :to="curentusername + '/' + curentuserid + '/' + user.Username + '/' + user.ID + '/chatroom'"
                                     class="user__tochatbtn">
                            チャットルームへ
                        </router-link>
                    </td>
                </tr>
            </table>

        </div>
    </div>
</template>

<script>
import Chatheader from "../components/Chatheader";
import axios from "axios";

export default {
    data() {
        return {
            users: [],
            curentuserid: localStorage.getItem('userID'),
            curentusername: localStorage.getItem('username')
        }
    },
    components: {
        Chatheader: Chatheader
    },
    created() {
        const params = new URLSearchParams();
        params.append('id',this.curentuserid);

        axios.post('/getusers', params).then(respons => {
            console.log(respons.data);
            this.users = respons.data;
        }).catch(error => {
            console.log(error);
        })
    },
    methods: {
        userchatroom(id) {
            const params = new URLSearchParams();
            params.append('curentuserid', this.curentuserid);
            params.append('seconduserid', id);
            axios.post('/userchatroom', params).then(response => {
                console.log(response);
            }).catch(error => {
                console.log(error);
            })
        }
    }
}
</script>

<style>

.users{
    width: 100%;
    padding: 40px 120px;
}

.users__ttl{
    font-size: 32px;
    display: inlin-block;
    margin: 0 auto;
    text-align: center;
}

.user__table {
    margin: 20px 100px;
    text-align: left;
    font-size: 18px;
}


.user__tochatbtn{
    padding: 0 20px;
    border: 2px solid #aaa;
    border-radius: 5px;
    margin-left: 20px;
}

</style>