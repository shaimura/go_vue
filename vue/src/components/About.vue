<template>
    <div id="about" class="about">
        <div class="about__conteainer">
            <h2 class="about__ttl">ABOUT</h2>
            <p class="about__subttl">- 自己紹介 -</p>
            <table class="about__table">
                <tr class="about__txt">
                    <th class="about__txt_ttl">名前：</th>
                    <td>志村 隼斗</td>
                </tr>
                <tr>
                    <th class="about__txt_ttl">経歴：</th>
                    <td>
                        2019年11月まで飲食店勤務し、<br>
                        その後2019年12月からIT企業の派遣社員として勤務
                    </td>
                </tr>
            </table>

        </div>

        <div class="about__img_container">
            <div class="about__img">
                <img src="../../src/assets/img/myimag025.jpg" alt="プロフィール画像" v-prlx="{ reverse: true, speed: 0.3 }">
            </div>
        </div>
    </div>
</template>

<script>

export default {

    methods: {
        parallax(){

        }
    }
}
</script>


<style>

.about{
    height: 100vh;
    width: 100%;
    background: radial-gradient( #D5EAD8 5%,#E9F4EB 45%,#ffffff 60%);
    padding: 80px 0px;
    display: flex;
}

.about__conteainer{
    margin: 180px 320px;
}

.about__ttl{
    font-size: 32px;
    letter-spacing: 0.2em;
    margin-bottom: 15px;
}

.about__subttl{
    font-size: 18px;
    letter-spacing: 0.2em;
    margin-bottom: 40px;
}

.about__table{
    width: 100%;
}

.about__txt_ttl{
    width: 20%;
}

.about__txt{
    letter-spacing: 0.1em;
    margin-bottom: 10px;
}

.about__img_container{
    position: relative;
    padding: 120px 0 0;
    height: 560px;
    width: 50%;

}

.about__img img{
    position: absolute;
    right: 0;
    height: 360px;
    width: 600px;
}




</style>