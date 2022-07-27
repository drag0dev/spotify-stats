export interface stats{
    href: string,
    items: string[],
    limit: number,
    next: string | null,
    offset: number,
    previous: string | null,
    total: number
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