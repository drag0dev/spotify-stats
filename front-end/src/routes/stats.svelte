<style>
    .stats{
        width: 60%;
        margin-left: 20%;
        margin-right: 20%;
    }

    .options{
        display: flex;
        flex-wrap: wrap;
        flex-direction: row;
        height: 80vh;
        max-height: 80vh;
    }

    .options button{
        width: 20%;
        height: 5%;

        color: #1DB954;
        background-color: black;
        border: 2px solid #1DB954;
        box-shadow: 3px 3px 3px 1px #1DB954;
        border-radius: 20px;
        cursor: pointer;
    }
    .options button:hover{
        border: 2px solid lightblue;
        box-shadow: 3px 3px 3px 1px lightblue;
        color: lightblue;
    }
</style>

<script>
    import { onMount } from 'svelte';

    import Header from "./_components/Header.svelte";
    import Footer from "./_components/Footer.svelte";

    import {grabCode, getStats, getAccessToken} from "../lib/stats";

    let message = '';
    let stats = false;

    onMount(async () => {
        let code = '';
        try{
            code = await grabCode();
        }catch(e){
            message = `Error: ${e}`;
        }
        message = 'Grabbing the stats...'

        let temp;
        try{
            temp = await grabCode();
        }catch(e){
            message = `Error: ${e}`;
        }

        let accessCode = '';
        try{
            accessCode = await getAccessToken(code);
        }catch(e){
            message = `Error: ${e}`;
        }
        console.log(accessCode)
        stats = true;
    });

</script>

<div class="stats">
    <Header />


    <div class="options">
        {#if stats === false}
            <p>{message}</p>
        {:else}
            <button>Artists</button>
            <button>Songs</button>
        {/if}
    </div>

    <Footer />
</div>