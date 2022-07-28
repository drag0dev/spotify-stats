<script type="ts">
    import type { track } from "$lib/interfaces";
    export let track: track;

</script>

<div class="track-div">
    <a href={track.external_urls.spotify} target="_blank"><img src={track.images[2].url} alt='artist'/></a>
    <div class="track-div-info">
        <a href={track.external_urls.spotify} target="_blank"><p class="track-name">{track.name}</p></a>
        <p>Release date: {track.release_date}</p>
        <p>Artists:
            {#each track.album.artists as artist}
                <a href={artist.external_urls.spotify} target="_blank">{artist.name} </a>
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
    }

    img{
        border: 2px solid #1DB954;
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

</style>