export default {
    // [min, max]
    Random(min, max) {
        return Math.round(Math.random() * (max - min)) + min;
    }
}