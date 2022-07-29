let SERVER_URL = "https://spotify-stats-production.up.railway.app/"
//let SERVER_URL = "http://localhost:8080/" //dev

import type { stats } from "./interfaces";

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

export const getStats = async (code: string): Promise<stats> => {
    let res = fetch(SERVER_URL + "stats", {
        method: "POST",
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({code: code})
    });
    if ((await res).status == 200){
        let data: Promise<stats> = (await res).json();
        return data;
    }
    return { // some form of error handling
        artists: [],
        tracks: []
    };
}
