<template>
    <div>
        <router-view name="Header"></router-view>
        <transition name="buttonfade">

        </transition>
        <router-view name="Top"></router-view>
        <router-view name="About"></router-view>
        <router-view name="Skill"></router-view>
        <router-view name="Portfolio"></router-view>
        <router-view name="Contact"></router-view>
        <router-view name="Footer"></router-view>
    </div>
</template>


<script>
import router from '../router';
export default {
    router
}
</script>

