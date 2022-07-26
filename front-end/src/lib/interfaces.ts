export interface stats{
    href: string,
    items: string[],
    limit: number,
    next: string | null,
    offset: number,
    previous: string | null,
    total: number
}