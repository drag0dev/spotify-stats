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

<script type="ts">
    import { onMount } from 'svelte';
    import type { artist } from '$lib/interfaces';

    import Header from "./_components/Header.svelte";
    import Footer from "./_components/Footer.svelte";

    import { grabCode, getStats } from "../lib/stats";
    import Artist from './_components/Artist.svelte';

    let message = '';
    let stats = false;
    let data: artist[];

    onMount(async () => {
        let code = '';
        try{
            code = await grabCode();
        }catch(e){
            message = `Error: ${e}`;
        }
        message = 'Grabbing the stats...'
        data = await getStats(code);
        stats = true;
    });

</script>

<div class="stats">
    <Header />


    <div class="options">
        {#if stats === false}
            <p>{message}</p>
        {:else}
            {#each data as a}
                <Artist artist={{...a, genres: a.genres.splice(0, 2)}}/>
            {/each}
        {/if}
    </div>

    <Footer />
</div>