function minTimeToType(word: string): number {
    if (word.length <= 0) {
        return 0;
    }

    let time = word.length;
    let prev = 'a'
    for (const cur of word) {
        time += circularDistance(cur, prev);
        prev = cur;
    }

    return time;
};

function circularDistance(cur: string, prev: string): number {
    const forwardDirection = Math.abs(cur.charCodeAt(0) - prev.charCodeAt(0));
    const reverseDirection = 26 - forwardDirection;
    return Math.min(forwardDirection, reverseDirection);

    // 2 - 26 = 24, 26 - 24 = 2, 1
}
