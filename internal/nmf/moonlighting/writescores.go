package moonlighting

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/knightjdr/cmgo/pkg/fs"
	"github.com/spf13/afero"
)

func writeMoonlightingScores(moonlightingScores moonScores, options outputOptions) {
	var buffer bytes.Buffer

	writeHeader(&buffer)
	writeBody(&buffer, moonlightingScores, options)

	afero.WriteFile(fs.Instance, options.outfile, buffer.Bytes(), 0644)
}

func writeHeader(buffer *bytes.Buffer) {
	buffer.WriteString("prey\t1st rank name\t2nd rank name\tmoonlighting score\t1st rank\t1st score\t2nd rank\t2nd score\n")
}

func writeBody(buffer *bytes.Buffer, moonlightingScores moonScores, options outputOptions) {
	for i, info := range moonlightingScores {
		gene := options.preyNames[i]
		primaryRank := fmt.Sprintf("%d", info.PrimaryRank+1)
		secondaryRank, secondaryScore, moonlightingScore := formatSecondaryRankAndMoonlightingScore(info, options.minRankValue)
		primaryRankName := defineRankName(primaryRank, options)
		secondaryRankName := defineRankName(secondaryRank, options)
		buffer.WriteString(
			fmt.Sprintf(
				"%s\t%s\t%s\t%s\t%s\t%0.5f\t%s\t%s\n",
				gene,
				primaryRankName,
				secondaryRankName,
				moonlightingScore,
				primaryRank,
				info.PrimaryScore,
				secondaryRank,
				secondaryScore,
			),
		)
	}
}

func formatSecondaryRankAndMoonlightingScore(info *preyInfo, minRankValue float64) (string, string, string) {
	if info.SecondaryRank == -1 || info.SecondaryScore < minRankValue {
		return "", "", ""
	}

	secondaryRank := fmt.Sprintf("%d", info.SecondaryRank+1)
	secondaryScore := fmt.Sprintf("%0.5f", info.SecondaryScore)
	moonlightingScore := fmt.Sprintf("%0.3f", info.MoonlightingScore)
	return secondaryRank, secondaryScore, moonlightingScore
}

func defineRankName(rank string, options outputOptions) string {
	if rank == "" {
		return ""
	}

	rankInt, _ := strconv.Atoi(rank)
	return strings.Join(options.localization[rankInt].DisplayTerms, ", ")
}
