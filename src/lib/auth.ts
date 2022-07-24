const clientId = 'f5940d4d679948c5a33bfce4ad03ac50';
const redirect_uri = 'http://localhost:5173/stats';
const scopes = 'user-read-recently-played user-top-read';
const baseUrl = 'https://accounts.spotify.com/authorize';

const auth = () => {
    window.location.replace(`${baseUrl}?response_type=code&client_id=${clientId}&redirect_uri=${redirect_uri}&scopes=${scopes}`);
}

export default auth;