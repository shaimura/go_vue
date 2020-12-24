<template>
    <div class="top">
        <div class="top__balloon" @click="scrollToSkill">
            <div class="top__container">
                <h2 class="top__tlt">WELCOME</h2>
                <p class="top__txt">ようこそ</p>
            </div>
        </div>

        <div class="arrow"></div>
        

        <div class="bubbles">
            <div class="bubble"></div>
            <div class="bubble"></div>
            <div class="bubble"></div>
            <div class="bubble"></div>
            <div class="bubble"></div>
        </div>

    </div>
</template>

<script>
export default {
    methods: {
        scrollToSkill() {
            const target = this.$el.getBoundingClientRect().bottom;
            window.scrollTo({
                top: target,
                behavior: "smooth"
            })
        }
    }
}
</script>


<style>

.top{
    height: 100vh;
    background: linear-gradient(to top, #ffffff,#D5EAD8);
    position: relative;
}

.top__balloon{
    height: 500px;
    width: 500px;
    background: radial-gradient( #BEDFC2 50%,#E9F4EB,#ffffff);
    border-radius: 50%;
    margin: auto;
    position: relative;
    top: 50%;
    transform: translateY(-50%);
    z-index: 10;
    cursor: pointer;
}

.top__balloon::before {
    content: '';
    display: inline-block;
    position: absolute;
    width: 100%;
    height: 100%;
    border-radius: 50%;
    top: 0;
    left: 0;
}


.top__balloon:hover::before{
    animation: sonar-effect 1.5s ease-in-out .3s infinite;
}

@keyframes sonar-effect {
  0% {
    opacity: 0.3;
  }
  40% {
    opacity: 0.5;
    box-shadow: 0 0 0 5px #fff, 0 0 10px 10px #fff, 0 0 0 10px #fff;
  }
  100% {
    box-shadow: 0 0 0 5px #fff, 0 0 10px 10px #fff, 0 0 0 10px #fff;
    transform: scale(1.5);
    opacity: 0;
  }
}




.top__container{
    text-align: center;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
}

.top__tlt{
    font-size: 48px;
    letter-spacing: 0.2em;
    margin-bottom: 15px;
    /* font-family: 'Roboto', sans-serif; */
}

.top__txt{
    font-size: 24px;
    letter-spacing: 0.2em;
}


.bubbles {
position: absolute;
width: 100%;
height: 100%;
z-index: 0;
overflow: hidden;
top: 0;
left: 0;
}
 
.bubble {
position: absolute;
bottom: -40;
width: 40px;
height: 40px;
background: radial-gradient( #ffffff,#E9F4EB);
border-radius: 50%;
animation: bubble 10s ease-in infinite;
}
 
.bubble:nth-child(1) {
width: 20px;
height: 20px;
left: 30%;
animation-duration: 6s;
}
 
.bubble:nth-child(2) {
left: 20%;
animation-duration: 3.5s;
}
 
.bubble:nth-child(3) {
width: 30px;
height: 30px;
left: 30%;
animation-duration: 7s;
}
 
.bubble:nth-child(4) {
width: 50px;
height: 50px;
left: 67%;
animation-duration: 6s;
}
 
.bubble:nth-child(5) {
width: 20px;
height: 20px;
left: 70%;
animation-duration: 4.5s;
}
 
@keyframes bubble {
  0% {
    bottom: -100px;
    transform: translateX(0);
  }
  50% {
    transform: translateX(100px);
  }
  100% {
    bottom: 1000px;
  }
}


.arrow{
    position: absolute;
    width: 40px;
    border-top: 1px solid #3EB370;
    transform: rotate(90deg);
    bottom: 50px;
    animation: fade-down 2s infinite;
}
.arrow::after{
    content: "";
    width: 40px;
    height: 15px;
    border-right: 1px solid #3EB370;
    transform: rotate(-45deg);
    position: absolute;
}

@keyframes fade-down {
  0% {
    opacity: 1;
    transform: rotate(90deg) translateX(-20px);
  }
  50% {
    opacity: 0.8;
    transform: rotate(90deg) translateX(15px);
  }
  100% {
    opacity: 0;
    transform: rotate(90deg) translateX(20px);
  }
}

</style>