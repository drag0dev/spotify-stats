export interface stats{
    artists: artist[],
    tracks: track[]
}

export interface artist{
    external_urls: {
        spotify: string
    },
    followers: {
        total: number
    },
    genres: string[],
    images: {
        height: number,
        width: number,
        url: string
    }[],
    name: string,
    popularity: number
}

export interface track{
    album: {
        artists: {
            name: string,
            external_urls: {
                spotify: string
            }
        }[]
    },
    external_urls: {
        spotify: string
    },
    images: {
        height: number,
        width: number,
        url: string
    }[],
    name: string,
    release_date: string
}