<style>
    .footer{
        width: 60vw;
        margin-left: 20vw;
        margin-right: 20vw;
        height: 10vh;
        display: flex;
        flex-wrap: wrap;
        flex-direction: row;
        justify-content: center;
        align-items: center;

        border-top: 2px solid #1DB954;
        font-weight: bold;
    }
    .footer p{
        text-align: center;
    }

    a{
        text-decoration: none;
        color: #1DB954;
    }
    @media screen and (max-width: 768px){
        .footer{
            width: 80vw;
            margin-left: 10vw;
            margin-right: 10vw;
        }

        .footer p{
            font-size: 3vw;
        }
    }
    @media screen and (max-width: 480px){
        .footer p{
            font-size: 4vw;
        }
    }
</style>


<div class="footer">
    <p>
        made by <a href="https://www.github.com/drag0dev" target="_blank">drag0dev</a> 2022 &copy;
    </p>
</div>