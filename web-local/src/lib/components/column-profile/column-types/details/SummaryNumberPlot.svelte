<script lang="ts">
  import {
    GraphicContext,
    SimpleDataGraphic,
  } from "@rilldata/web-common/components/data-graphic/elements";
  import { DynamicallyPlacedLabel } from "@rilldata/web-common/components/data-graphic/guides";
  import { INTEGERS } from "@rilldata/web-common/lib/duckdb-data-types";
  import { justEnoughPrecision } from "@rilldata/web-common/lib/formatters";
  import { format } from "d3-format";

  export let min;
  export let max;
  export let q25;
  export let q50;
  export let q75;
  export let mean;
  export let rowHeight = 24;
  export let type: string;

  $: formatter = INTEGERS.has(type) ? format(".0r") : justEnoughPrecision;
  $: values = [
    { label: "min", value: min, format: formatter },
    { label: "q25", value: q25, format: formatter },
    { label: "q50", value: q50, format: formatter },
    {
      label: "q75",
      value: q75,
      format: formatter,
    },

    { label: "max", value: max, format: formatter },
    {
      label: "mean",
      value: mean,
      format: justEnoughPrecision,
    },
  ].reverse();
</script>

{#if values}
  <!-- note: this currently inherits its settings from a parent GraphicContext -->
  <SimpleDataGraphic
    top={0}
    buffer={0}
    height={values?.length * rowHeight}
    let:xScale
  >
    {#each values as { label, value, format = undefined }, i}
      <g transform="translate(0 {(values.length - i - 1) * rowHeight})">
        <GraphicContext height={rowHeight}>
          <circle
            cx={xScale(value)}
            cy={rowHeight / 2}
            r="2.5"
            fill="#ff8282"
          />
          <line
            x1={xScale(value)}
            x2={xScale(value)}
            y1={rowHeight / 2 - 1}
            y2={-(rowHeight * (values.length - i - 1))}
            stroke="red"
            opacity={0.5}
          />
          <DynamicallyPlacedLabel
            dy=".35em"
            x={value}
            ry={rowHeight / 2}
            buffer={6}
            colorClass="ui-copy-muted"
          >
            <tspan>
              <tspan class="font-semibold ui-copy-muted">{label}</tspan>
              <tspan class="ui-copy-disabled"
                >{format ? format(value) : value}</tspan
              >
            </tspan>
          </DynamicallyPlacedLabel>
        </GraphicContext>
      </g>
    {/each}
  </SimpleDataGraphic>
{/if}
