<template>
    <div id="contact" class="contact">
        <div class="contact__top">
            <div class="contact__position_set">
                <h2 class="contact__ttl">CONTACT</h2>
                <p class="contact__subttl">- お問い合わせ -</p>
            </div>
        </div>
        
        <div class="contact__bottom">
            <div class="contact__position_set">
                <p class="contact__txt">お気軽にこちらのアドレスまでご連絡ください。</p>
                <a href="mailto:daken240515@gmail.com"><p class="contact__address">daken240515@gmail.com</p></a>
            </div>
        </div>
    </div>
</template>


<style>

.contact{
    height: 90vh;
    width: 100%;
    padding: 10vh 0 0;
}

.contact__top{
    width: 50%;
    height: 35vh;
    margin: auto;
    background: radial-gradient( #D5EAD8 5%,#E9F4EB 45%,#ffffff 60%);
    border-radius: 50%;
    position: relative;
}

.contact__position_set{
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    text-align: center;
}

.contact__ttl{
    font-size: 32px;
    letter-spacing: 0.2em;
    margin-bottom: 15px;
}

.conatct__subttl{
    font-size: 18px;
    letter-spacing: 0.2em;
}

.contact__bottom{
    width: 90%;
    height: 40vh;
    margin: auto;
    background: radial-gradient( #D5EAD8 5%,#E9F4EB 45%,#ffffff 60%);
    border-radius: 50%;
    text-align: center;
    position: relative;
}

.contact__txt{
    font-size: 18px;
    letter-spacing: 0.1em;
    margin-bottom: 15px;
}

.contact__address{
    font-size: 24px;
    letter-spacing: 0.2em;
}

</style>