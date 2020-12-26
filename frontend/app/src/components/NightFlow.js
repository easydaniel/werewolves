import React, { useState } from "react"

import {
  Stepper,
  StepLabel,
  StepContent,
  Step,
  Button,
  Typography,
} from "@material-ui/core"

const NightFlow = ({ flow }) => {
  const [activeStep, setActiveStep] = useState(0)

  return (
    flow && (
      <>
        <Stepper activeStep={activeStep} orientation="vertical">
          {flow.map((label, idx) => (
            <Step key={idx}>
              <StepLabel>
                <Typography variant="h6" component="h1">
                  {label}
                </Typography>
              </StepLabel>
              <StepContent>
                <Button
                  disabled={activeStep === 0}
                  onClick={() => setActiveStep(prevStep => prevStep - 1)}
                >
                  上一步
                </Button>
                <Button
                  variant="contained"
                  color="primary"
                  onClick={() =>
                    activeStep === flow.length - 1
                      ? setActiveStep(0)
                      : setActiveStep(prevStep => prevStep + 1)
                  }
                >
                  {activeStep === flow.length - 1 ? "重新開始" : "下一步"}
                </Button>
              </StepContent>
            </Step>
          ))}
        </Stepper>
      </>
    )
  )
}

export default NightFlow
