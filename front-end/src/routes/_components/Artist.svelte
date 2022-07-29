<script type="ts">
    import type { artist } from "$lib/interfaces";
    export let artist: artist;
</script>


<div class="artist-div">
    <a href={artist.external_urls.spotify} target="_blank"><img src={artist.images[2].url} alt='artist'/></a>
    <div class="artist-div-info">
        <a href={artist.external_urls.spotify} target="_blank"><p class="artist-name">{artist.name}</p></a>
        <p>Followers: {artist.followers.total}</p>
        <p>Popularity: {artist.popularity}/100</p>
        <p>Genres: 
            {#each artist.genres as g, i (i)}
                {g + " "} 
            {/each}
        </p>
    </div>
</div>

<style>
    .artist-div{
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
    .artist-div-info{
        width: 50%;
        margin-left: 1vw;
    }

    .artist-name{
        color: #1DB954;
    }
    
    a{
        text-decoration: none;
    }

</style>