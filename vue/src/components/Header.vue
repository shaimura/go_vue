<template>
    <header class="header" :class="{header__background: addBackground}">
        <p>My Portfolio</p>

        <nav class="header__nav">
            <ul id="nav" class="header__list">
                <li data-menuanchor="about" class="header__list_item active"><a href="#top">TOP</a></li>
                <li data-menuanchor="about" class="header__list_item"><a href="#about">ABOUT</a></li>
                <li data-menuanchor="skill" class="header__list_item"><a href="#skill">SKILL</a></li>
                <li data-menuanchor="portfolio" class="header__list_item"><a href="#portfolio">PORTFOLIO</a></li>
                <li data-menuanchor="contact" class="header__list_item"><a href="#contact">CONTACT</a></li>
            </ul>
        </nav>
    </header>
</template>


<script>
export default {
    data() {
        return {
            addBackground: false
        }
    },
    created() {
        window.addEventListener("scroll", this.addBackgroundCllor)
    },
    distroyed() {
        window.removeEventListener("scroll", this.addBackgroundCllor)
    },
    methods: {
        addBackgroundCllor() {
            if(window.pageYOffset > window.innerHeight/2){
                this.addBackground = true
            } else {
                this.addBackground = false
            }
        }
    }
}
</script>


<style>
.header{
    display: flex;
    justify-content: space-between;
    align-items: center;
    padding: 0 120px;
    background: rgba(235, 245, 236, 0);
    height: 10vh;
    width: 100%;
    z-index: 99;
    position: fixed;
    transition: 1s;
}

.header__background{
    background: rgba(235, 245, 236, 0.8);
    transition: 1s;
}

.header__list{
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.header__list_item{
    margin-left: 20px;
    letter-spacing: 0.1em;
}


</style>