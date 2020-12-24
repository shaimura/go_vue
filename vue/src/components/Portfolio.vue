<template>
    <div id="portfolio" class="portfolio">
        <div class="portfolio__backgrounds">
            <div class="portfolio__background"></div>
            <div class="portfolio__background"></div>
            <div class="portfolio__background"></div>
        </div>

        <div class="portfolio__container">
            <h2 class="portfolio__ttl">PORTFOLIO</h2>
            <p class="portfolio__subttl">- 実績 -</p>


            <div class="portfolio__imags_2">
                <div class="portfolio__img_2">
                    <img src="../../src/assets/img/portfollio.png" alt="サンプル">
                    <p class="portfolio__img_txt">GOとVue.jsを使用してチャット機能を作成しました。</p>
                    <router-link :to="'/sinup'" class="portfolio__img_txt">制作物へ</router-link>
                </div>
            </div>
        </div>
    </div>
</template>


<style>

.portfolio{
    height: 100vh;
    width: 100%;
    padding: 10vh 0;
    text-align: center;
    position: relative;
}

.portfolio__backgrounds{
    display: flex;
    position: absolute;
    z-index: -99;
    height: 100vh;
    width: 100%;
    overflow: hidden;
}

.portfolio__background{
    height: 100vh;
    width: 50%;
    background: radial-gradient( #D5EAD8 5%,#E9F4EB 45%,#ffffff 60%);
    position: absolute;
    border-radius: 50%;
}


.portfolio__background:nth-child(1){
    left: -50%;
    transform: translateX(50%);


}

.portfolio__background:nth-child(2){
    left: 50%;
    transform: translateX(-50%);
}

.portfolio__background:nth-child(3){
    right: -50%;
    transform: translateX(-50%);

}

.portfolio__ttl{
    font-size: 32px;
    letter-spacing: 0.2em;
    margin-bottom: 15px;
    padding-top: 160px;
}

.portfolio__subttl{
    font-size: 18px;
    letter-spacing: 0.2em;
    margin-bottom: 40px;
}

.portfolio__imags{
    display: flex;
}

.portfolio__img_txt{
    display: block;
    margin-top: 10px;
}


/* test */
.portfolio__imags_2 {
    width: 60%;
    height: 60%;
    margin: auto;
}

.portfolio__img_2 img{
    width: 60%;
    height: 60%;
}

</style>