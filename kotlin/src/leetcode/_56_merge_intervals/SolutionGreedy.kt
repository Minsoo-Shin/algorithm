package leetcode._56_merge_intervals

class SolutionGreedy {
    fun merge(intervals: Array<IntArray>): Array<IntArray> {
        if (intervals.size <= 1) {
            return intervals
        }
        // end 지점 저장
        val result = mutableListOf<IntArray>()

        val sorted = intervals.sortedBy { it[0] }

        var start = sorted[0][0]
        var end = sorted[0][1]

        for (i in 1 until intervals.size) {
            val (sessionStart, sessionEnd) = sorted[i]

            if (end < sessionStart) {
                result.add(intArrayOf(start, end))
                start = sessionStart
                end = sessionEnd
            } else {
                end = maxOf(end, sessionEnd)
            }
        }
        result.add(intArrayOf(start, end))

        return result.toTypedArray()
    }
}