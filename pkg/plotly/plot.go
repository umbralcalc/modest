package plotly

import (
	"fmt"
	"strings"
)

type Plot struct {
	Dates   []string
	Amounts []string
}

func (p Plot) String() string {
	return fmt.Sprintf(`
	<div id="plot" style="width:100%%;height:100%%;"></div>
	<script>
		var trace = {
			x: [%s],
			y: [%s],
			mode: 'markers',
			type: 'scatter',
			marker: { size: 10 },
			name: 'Amount Over Time'
		};
		var data = [trace];
		var layout = {
			title: 'Amount Over Time',
			xaxis: { title: 'Date' },
			yaxis: { title: 'Amount' }
		};
		Plotly.newPlot('plot', data, layout);
	</script>`, strings.Join(p.Dates, ","), strings.Join(p.Amounts, ","))
}
