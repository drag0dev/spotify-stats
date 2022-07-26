const clientId = 'f5940d4d679948c5a33bfce4ad03ac50';
const redirect_uri = 'https://spotify-stats-gray.vercel.app/stats';
// const redirect_uri = 'http://localhost:5173/stats'; //dev
const scopes = 'user-top-read';
const baseUrl = 'https://accounts.spotify.com/authorize';

const auth = () => {
    window.location.href = `${baseUrl}?response_type=code&client_id=${clientId}&redirect_uri=${redirect_uri}&scope=${scopes}`;
}

export default auth;