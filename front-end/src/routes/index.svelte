<script>
    import auth from "../lib/auth";
</script>

<div class="index">
    <div class="index-message">
        <p>
            Want to see your top songs and artists? Just click Log in!
        </p>
    </div>

    <div class="options" on:click={auth}>
            <button>Log in</button>
    </div>
</div>

<style>
    *{
        margin: 0;
    }
    .index{
        display: flex;
        flex-wrap: wrap;
        flex-direction: row;
        margin-left: 20vw;
        margin-right: 20vw;
        width: 60vw;
    }

    .index-message{
        display: flex;
        justify-content: center;
        align-items: center;
        text-align: center;
        width: 100%;
        padding-top: 2vh;
    }

    .index-message p{
        text-align: center;
        color: #1DB954;
        font-weight: bold;
    }
    .options{
        width: 100%;
        display: flex;
        flex-wrap: wrap;
        flex-direction: row;
        justify-items: center;
        justify-content: center;
        align-items: center;
        height: 15vh;
    }
    .options button{
        width: 20%;
        height: 25%;

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

    @media screen and (max-width: 768px){
        .index{
            margin-left: 10vw;
            margin-right: 10vw;
            width: 80vw;
        }
        p{
            font-size: 3vw;
        }

        .options button{
            width: 25%;
            font-size: 3vw;
        }
    }
    @media screen and (max-width: 480px){
        .index p{
            font-size: 4vw;
        }
        .options button{
            width: 30%;
            height: 4vh;
            font-size: 4vw;
        }
    }
</style>