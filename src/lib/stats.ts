const baseUrlAT = 'https://accounts.spotify.com/api/token'
const baseUrlTop = 'https://api.spotify.com/v1/me/top/type';
const clientId = 'f5940d4d679948c5a33bfce4ad03ac50';
const redirect_uri = 'http://localhost:5173/stats';

export const grabCode = (): Promise<string> => {
   let res: Promise<string> = new Promise((reslove, reject) => {
        let urlParams = new URLSearchParams(window.location.search);
        if (!urlParams.has('code')){
            reject('invalid url');
        }else{
            reslove(urlParams.get('code')!);
        }
   }); 
   return res;
}

export const getAccessToken = (code: string): Promise<string> => {
    let promise: Promise<string> = new Promise(async (resolve, reject) => {
        let res = await fetch(baseUrlAT, {
            method: "POST",
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded',
                'Authorization': 'Basic ' + btoa(`${clientId}:`)
            },
            body: new URLSearchParams({
                'grant_type': 'authorization_code',
                'code': code,
                'redirect_uri': 'http://localhost:5173/stats'
            })
        });
        let data = await res.json();
        if (data.error){
            reject(data.error);
        } else{
            resolve (data.access_token);
        }
    });
    return promise;
}

export const getStats = () => {

}
