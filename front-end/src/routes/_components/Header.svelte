<style>
    .header{
        width: 60vw;
        margin-left: 20vw;
        margin-right: 20vw;

        height: 10vh;
        max-height: 10vh;
        text-align: center;

        display: flex;
        justify-content: center;
        align-items: center;

        border-bottom: 2px solid #1DB954;
    }
    .header h1{
        font-size: 4vw;
    }

    @media screen and (max-width: 768px){
        .header{
            width: 80vw;
            margin-left: 10vw;
            margin-right: 10vw;
        }

        .header h1{
            font-size: 7vw;
        }
    }

    @media screen and (max-width: 320px){
        .header h1{
            font-size: 10vw;
        }
    }
</style>


<div class="header">
    <h1>Spotify Stats</h1>
</div>
