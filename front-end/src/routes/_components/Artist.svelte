<script type="ts">
    import type { artist } from "$lib/interfaces";
    export let artist: artist;
    export let index: number;
    let len: number;
    $: len = artist.genres.length
</script>


<div class="artist-div {index % 2 == 0 ? 'even' : 'odd'}">
    <div class="index">
        <p class="index">{index+1}.</p>
    </div>
    <a href={artist.external_urls.spotify} target="_blank"><img src={artist.images[2].url} alt='artist'/></a>
    <div class="artist-div-info">
        <a href={artist.external_urls.spotify} target="_blank"><p class="artist-name">{artist.name}</p></a>
        <p>Followers: {artist.followers.total}</p>
        <p>Popularity: {artist.popularity}/100</p>
        <p>Genres: 
            {#each artist.genres as g, i (i)}
                {g + " "} {i + 1 != len ? ' | ' : ''} 
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

    .even{
        background-color: rgba(0, 0, 0, 0.1);
    }

</style>