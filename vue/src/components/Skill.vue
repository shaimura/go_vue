<template>
        <div id="skill" class="skill">
            <div class="skill__container">
                    <h2 class="skill__ttl">SKILL</h2>
                    <p class="skill__subttl">- スキル -</p>
                <transition name="fade">

                    <div v-show="visibleSkill">

                        <div class="skill__list">
                            <ul class="skill__list_left">
                                <li class="skill__list_item">HTML,CSS</li>
                                <li class="skill__list_item">WordPress</li>
                                <li class="skill__list_item">Golang</li>
                            </ul>

                            <ul class="skill__list_rigth">
                                <li class="skill__list_item">JavaScript</li>
                                <li class="skill__list_item">Vue.js</li>
                                <li class="skill__list_item">Docker</li>
                            </ul>
                        </div>
                    </div>

            </transition>
            </div>

            <transition name="fade">
                <div v-show="visibleSkill">

                    <div class="skill__explanation">
                        <p class="skill__explanation_txt">
                            スクールでHTML・CSS、JavaScript、WordPressを学び、独学でVue.js、Golang、Dockerを勉強しています。
                        </p>
                    </div>
                </div>
            </transition>
        </div>
</template>


<script>

export default {
    data() {
        return {
            visibleSkill: false
        }
    },
    created() {
        window.addEventListener("scroll", this.handleScroll)
    },
    distroyed() {
        window.removeEventListener("scroll", this.handleScroll)
    },
    methods: {
        handleScroll() {
            if(!this.visibleSkill) {
                var top = this.$el.getBoundingClientRect().top;
                this.visibleSkill = top < window.innerHeight.valueOf()/2;
            }
        }
    }
}
</script>

<style>

.skill{
    height: 100vh;
    width: 100%;
    padding: 10vh 0;
    display: flex;
    position: relative;
}

.skill__container{
    height: 60vh;
    width: 50%;
    background: radial-gradient( #D5EAD8 5%,#E9F4EB 45%,#ffffff 60%);
    padding: 120px 240px;
    border-radius: 50%;
}

.skill__ttl{
    font-size: 32px;
    letter-spacing: 0.2em;
    margin-bottom: 15px;
}

.skill__subttl{
    font-size: 18px;
    letter-spacing: 0.2em;
    margin-bottom: 40px;
}

.skill__list{
    display: flex;

}

.skill__list_item{
    margin: 0 15px 10px 0;
    font-size: 16px;
    letter-spacing: 0.1em;
}

.skill__list_left{
    margin-right: 20px;
}


.skill__explanation{
    height: 60vh;
    width: 50%;
    background: radial-gradient( #D5EAD8 5%,#E9F4EB 45%,#ffffff 60%);
    padding: 160px 200px;
    border-radius: 50%;
    position: absolute;
    bottom: 0;
    right: 0;
}


.skill__explanation_txt{
    font-size: 16px;
    letter-spacing: 0.2em;
    line-height: 1.6em;
}


/* スライドイン */
.fade-enter-active, .fade-leave-active{
    transition: opacity 3s
}

.fade-enter, .fade-leave-to{
    opacity: 0;
}

</style>

