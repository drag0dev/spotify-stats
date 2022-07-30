<script type="ts">
    import type { track } from "$lib/interfaces";
    export let track: track;
    export let index: number;
    let len: number;
    $: len = track.album.artists.length
</script>

<div class="track-div {index % 2 == 0 ? 'even' : 'odd'}">
    <div class="index">
        <p class="index">{index+1}.</p>
    </div>
    <a href={track.external_urls.spotify} target="_blank"><img src={track.album.images[1].url} alt='artist'/></a>
    <div class="track-div-info">
        <a href={track.external_urls.spotify} target="_blank"><p class="track-name ">{track.name}</p></a>
        <p>Release date: {track.album.release_date}</p>
        <p>Popularity: {track.popularity}/100</p>
        <p>Artists:
            {#each track.album.artists as artist, i (i)}
                <a href={artist.external_urls.spotify} target="_blank" class="artist-name">{artist.name}</a>{i+1 != len ? ' | ' : ''} 
            {/each}
        </p>
    </div>
</div>


<style>
    .track-div{
        display: flex;
        flex-direction: row;
        flex-wrap: wrap;
        margin-top: 2vh;
        margin-bottom: 2vh;
        width: 100%;
        justify-content: center;
        padding-top: 2vh;
        padding-bottom: 2vh;
        border-radius: 20px;
    }

    .index{
        display: flex;
        justify-content: center;
        align-items: center;
        margin-right: 2vw;
        font-size: 1.5vw;
        color: #1DB954;
    }
    img{
        border: 2px solid #1DB954;
        max-width: 160px;
        max-height: 160px;
    }
    .track-div-info{
        width: 50%;
        margin-left: 1vw;
    }
    .track-name{
        color: #1DB954;
    }
    a{
        text-decoration: none;
    }
    .artist-name{
        color: #1DB954;
    }

    .even{
        background-color: rgba(0, 0, 0, 0.1);
    }

</style>