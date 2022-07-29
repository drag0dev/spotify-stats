<script type="ts">
    import { onMount } from 'svelte';
    import Header from "./_components/Header.svelte";
    import Footer from "./_components/Footer.svelte";
    import Artist from './_components/Artist.svelte';
    import Track from './_components/Track.svelte';

    import type { stats } from '$lib/interfaces';
    import { grabCode, getStats } from "../lib/stats";

    enum Option{
        None,
        Tracks,
        Artists
    }

    let message = '';
    let stats = false;
    let data: stats;
    let option: Option = Option.None;

    onMount(async () => {
        let code = '';
        try{
            code = await grabCode();
        }catch(e){
            message = `Error: ${e}`;
            return;
        }

        message = 'Grabbing the stats...'
        data = await getStats(code);
        if (data.artists.length == 0 && data.tracks.length == 0){
            message = 'No stats to show!'
            return;
        }
        stats = true;
    });
</script>

<div class="stats">
    <Header />

    {#if stats === false}
        <p class="message-p">{message}</p>
    {:else}
        <div class="options">
            <button on:click={() => {option = Option.Tracks}}>Tracks</button>
            <button on:click={() => {option = Option.Artists}} class="artists-button">Artists</button>
        </div>
    {/if}

    {#if option == Option.Tracks}
        {#each data.tracks as t (t.name)}
            <Track track={t}/>
        {/each}
    {/if}

    {#if option == Option.Artists}
        {#each data.artists as a (a.name)}
            <Artist artist={a}/>
        {/each}
    {/if}

    <Footer />
</div>

<style>
    .stats{
        width: 60%;
        min-height: 80vh;
        margin-left: 20%;
        margin-right: 20%;
    }

    .options{
        display: flex;
        flex-wrap: wrap;
        flex-direction: row;
        justify-content: center;
        height: 20%;
        width: 100%;
        padding-top: 2vh;
        padding-bottom: 2vh;
    }

    .artists-button{
        margin-left: 2vw;
    }

    .options button{
        width: 15%;
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

    .message-p{
        text-align: center;
        color: red;
        font-size: 2vw;
    }
</style>
