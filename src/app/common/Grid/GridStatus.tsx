import React from "react";
import styled from "styled-components";

const SGridStatus = styled.div`
  grid-area: status;
`;

export interface GridStatusProps {
}

export class GridStatus extends React.Component<GridStatusProps> {
  render() {
    return (
      <SGridStatus>{this.props.children}</SGridStatus>
    );
  }
}