import React from "react"

import {
  Stepper,
  StepLabel,
  StepContent,
  Step,
  Typography,
} from "@material-ui/core"

const NightFlow = ({ flow }) => {
  return (
    flow && (
      <>
        <Stepper orientation="vertical">
          {flow.map((label, idx) => (
            <Step key={idx} completed={false} active>
              <StepLabel>
                <Typography variant="subtitle2">{label}</Typography>
              </StepLabel>
              <StepContent></StepContent>
            </Step>
          ))}
        </Stepper>
      </>
    )
  )
}

export default NightFlow
