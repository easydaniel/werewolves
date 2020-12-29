import React, { useState } from "react"

import {
  Dialog,
  DialogTitle,
  DialogActions,
  Button,
  DialogContent,
  FormLabel,
  FormControl,
  FormControlLabel,
  Radio,
  RadioGroup,
} from "@material-ui/core"

const SelectDialog = ({
  open,
  availableOptions,
  title,
  dismissDialog,
  submit,
}) => {
  const [value, setValue] = useState("")

  return (
    <Dialog open={open}>
      <DialogTitle>{title}</DialogTitle>
      <DialogContent>
        <FormControl component="fieldset">
          <RadioGroup value={value} onChange={e => setValue(e.target.value)}>
            {availableOptions.map(({ label, disabled }, idx) => (
              <FormControlLabel
                key={idx}
                value={`${idx}`}
                control={<Radio />}
                label={label}
                disabled={disabled}
              />
            ))}
          </RadioGroup>
        </FormControl>
      </DialogContent>
      <DialogActions>
        <Button onClick={() => dismissDialog()} color="secondary">
          取消
        </Button>
        <Button
          onClick={() => {
            submit(value)
            setValue("")
          }}
          color="primary"
          autoFocus
        >
          確認
        </Button>
      </DialogActions>
    </Dialog>
  )
}

export default SelectDialog
