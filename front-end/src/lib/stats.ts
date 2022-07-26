let SERVER_URL = "https://spotify-stats-production.up.railway.app/"

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

export const getStats = async (code: string): Promise<stats | undefined> => {
    let res = fetch(SERVER_URL + "stats", {
        method: "POST",
        headers: {
            'Content-Type': 'application/json:'
        },
        body: JSON.stringify({code: code, stat: 3}) // TODO: different stats
    });
    if ((await res).status != 200){
        console.log('cant get stats'); //TODO: only for now
    }else{
        let data: Promise<stats> = (await res).json();
        return data;
    }
    return undefined;
}
